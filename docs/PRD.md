# Product Requirement Document (PRD)

## 1. Ringkasan Produk
**Mini Bank** adalah layanan sistem perbankan berkonsep inti (core banking/wallet system) berskala ringan yang melayani siklus transaksi finansial seperti pendaftaran rekening, penyimpanan dana, penarikan, transfer antar-rekening, serta pencatatan mutasi transaksi berseri (immutable transaction history).

## 2. Target Pengguna
- **Nasabah (Customer):** Melakukan pengecekan saldo, mutasi, deposit, withdraw, dan transfer.
- **Admin/Teller:** Mengelola akun nasabah dan rekonsiliasi kas internal (jika diperlukan).

## 3. Ruang Lingkup Fitur (Core Features)
1. **Manajemen Rekening (Account Management):**
   - Pembuatan rekening bank (Account ID, Nama, Nomor Rekening, Saldo Awal).
   - Penarikan profil rekening dan validasi status rekening (Aktif/Diblock).
2. **Transaksi Finansial (Financial Transactions):**
   - **Deposit:** Setor tunai ke dalam rekening.
   - **Withdraw:** Tarik tunai dari rekening dengan pengecekan saldo minimum.
   - **Transfer:** Pemindahan dana antar dua rekening secara atomik (ACID transactions).
3. **Riwayat Buku Besar (Ledger & Audit Trail):**
   - Setiap perubahan saldo direkam dalam tabel transaksi & mutasi ledger yang konklusif dan tidak dapat di-edit/delete (Immutable).
