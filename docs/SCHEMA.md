# Skema Database & Relasi Tabel (SCHEMA)

Penyimpanan utama aplikasi **Mini Bank** mengandalkan **PostgreSQL** untuk data transaksi finansial relasional dan **Redis** untuk manajemen sesi serta *rate limiting*. Skema dirancang mendukung integrasi Payment Gateway (*Midtrans / Xendit*) serta pencatatan audit (*Audit Logs & Journal Entries*).

## 1. Daftar Tabel Utama di PostgreSQL

### `users`
Menyimpan kredensial autentikasi baik Nasabah (Customer) maupun Administrator.
```sql
CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    fullname VARCHAR(150) NOT NULL,
    email VARCHAR(150) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    role VARCHAR(30) DEFAULT 'customer' NOT NULL, -- 'customer' atau 'admin'
    status VARCHAR(30) DEFAULT 'active' NOT NULL, -- 'active', 'suspended'
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_users_email ON users(email);
```

### `accounts`
Menyimpan informasi rekening nasabah dan saldo akhir kas.
```sql
CREATE TABLE IF NOT EXISTS accounts (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID REFERENCES users(id) ON DELETE RESTRICT,
    account_number VARCHAR(20) UNIQUE NOT NULL,
    balance NUMERIC(15, 2) NOT NULL DEFAULT 0.00 CHECK (balance >= 0),
    currency VARCHAR(3) NOT NULL DEFAULT 'IDR',
    status VARCHAR(20) DEFAULT 'active' NOT NULL, -- 'active', 'frozen'
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_accounts_user_id ON accounts(user_id);
CREATE INDEX idx_accounts_number ON accounts(account_number);
```

### `transactions`
Mencatat riwayat mutasi perpindahan uang baik Internal Transfer, Deposit, maupun Withdraw.
```sql
CREATE TYPE transaction_type AS ENUM ('DEPOSIT', 'WITHDRAW', 'TRANSFER', 'TOPUP_PAYOUT');

CREATE TABLE IF NOT EXISTS transactions (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    reference_no VARCHAR(50) UNIQUE NOT NULL,
    from_account_id UUID REFERENCES accounts(id) ON DELETE RESTRICT,
    to_account_id UUID REFERENCES accounts(id) ON DELETE RESTRICT,
    amount NUMERIC(15, 2) NOT NULL CHECK (amount > 0),
    tx_type transaction_type NOT NULL,
    status VARCHAR(30) DEFAULT 'SUCCESS' NOT NULL,
    description TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_transactions_ref ON transactions(reference_no);
CREATE INDEX idx_transactions_from ON transactions(from_account_id);
CREATE INDEX idx_transactions_to ON transactions(to_account_id);
```

### `journal_entries` (General Ledger)
Pencatatan buku besar akuntansi ganda (*Double-Entry Ledger*) untuk audit rekonsiliasi kas bank murni (Debit vs Kredit).
```sql
CREATE TYPE entry_type AS ENUM ('DEBIT', 'CREDIT');

CREATE TABLE IF NOT EXISTS journal_entries (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    transaction_id UUID REFERENCES transactions(id) ON DELETE RESTRICT,
    account_id UUID REFERENCES accounts(id) ON DELETE RESTRICT,
    amount NUMERIC(15, 2) NOT NULL CHECK (amount > 0),
    type entry_type NOT NULL,
    balance_before NUMERIC(15, 2) NOT NULL,
    balance_after NUMERIC(15, 2) NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_journal_tx ON journal_entries(transaction_id);
CREATE INDEX idx_journal_account ON journal_entries(account_id);
```

### `payment_orders` (Top Up & Payment Gateway Orders)
Mencatat tagihan pemesanan saldo melalui pihak eksternal (*Midtrans / Xendit*) dan verifikasi *webhook*.
```sql
CREATE TABLE IF NOT EXISTS payment_orders (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    order_id VARCHAR(50) UNIQUE NOT NULL,
    account_id UUID REFERENCES accounts(id) ON DELETE RESTRICT,
    payment_gateway VARCHAR(50) NOT NULL, -- 'midtrans' / 'xendit'
    amount NUMERIC(15, 2) NOT NULL CHECK (amount > 0),
    payment_method VARCHAR(50), -- e.g., 'VA_BCA', 'QRIS'
    status VARCHAR(30) DEFAULT 'PENDING' NOT NULL, -- 'PENDING', 'SETTLED', 'FAILED', 'EXPIRED'
    webhook_payload JSONB,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_payment_orders_id ON payment_orders(order_id);
```

### `audit_logs`
Rekam jejak aktivitas sensitif di sistem (login, penggantian sandi, pemblokiran akun, penyesuaian limit).
```sql
CREATE TABLE IF NOT EXISTS audit_logs (
    id BIGSERIAL PRIMARY KEY,
    user_id UUID,
    action VARCHAR(100) NOT NULL,
    entity_type VARCHAR(50) NOT NULL,
    entity_id VARCHAR(100),
    old_values JSONB,
    new_values JSONB,
    ip_address VARCHAR(45),
    user_agent TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_audit_user ON audit_logs(user_id);
CREATE INDEX idx_audit_action ON audit_logs(action);
```

---

## 2. Struktur Data Caching di Redis

1. **Session Store & Tokens:**
   - Key: `session:jwt:<user_id>:<token_id>` -> Val: metadata status (TTL: 24 jam).
2. **Rate Limiting:**
   - Key: `ratelimit:ip:<client_ip>` -> Val: hit count (TTL: 1 menit, Max: 60 req/min).
   - Key: `ratelimit:transfer:<account_id>` -> Proteksi spam transfer berurutan.
3. **Temporary Cache (Account Profile):**
   - Key: `cache:account:<account_number>` -> Val: JSON profil ringkas (TTL: 15 menit).
