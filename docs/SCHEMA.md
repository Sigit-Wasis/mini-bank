# Skema Database & Relasi Tabel (SCHEMA)

Database yang digunakan pada core mini-bank ini adalah **PostgreSQL**.

## 1. Daftar Tabel Utama (Entities)

### `accounts`
Menyimpan informasi identifikasi utama dan akumulasi saldo pengguna.
```sql
CREATE TABLE IF NOT EXISTS accounts (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    account_number VARCHAR(20) UNIQUE NOT NULL,
    owner_name VARCHAR(100) NOT NULL,
    balance NUMERIC(15, 2) NOT NULL DEFAULT 0.00 CHECK (balance >= 0),
    currency VARCHAR(3) NOT NULL DEFAULT 'IDR',
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_accounts_number ON accounts(account_number);
```

### `transactions`
Mencatat metadata dan riwayat seluruh jenis perpindahan dana (Deposit, Withdraw, Transfer).
```sql
CREATE TYPE transaction_type AS ENUM ('DEPOSIT', 'WITHDRAW', 'TRANSFER');

CREATE TABLE IF NOT EXISTS transactions (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    from_account_id UUID REFERENCES accounts(id) ON DELETE RESTRICT,
    to_account_id UUID REFERENCES accounts(id) ON DELETE RESTRICT,
    amount NUMERIC(15, 2) NOT NULL CHECK (amount > 0),
    tx_type transaction_type NOT NULL,
    description TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_transactions_from ON transactions(from_account_id);
CREATE INDEX idx_transactions_to ON transactions(to_account_id);
```

## 2. Aturan Relasional & Integritas (Constraints)
- **Zero/Positive Balance Lock:** `CHECK (balance >= 0)` mencegah anomali minus saat *concurrency bug* lepas di layer aplikasi.
- **Immutable Transactions:** Tidak ada hak akses `UPDATE` maupun `DELETE` ke tabel `transactions` setelah disuntikkan.
