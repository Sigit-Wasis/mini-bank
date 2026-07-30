# Roadmap Pengembangan (ROADMAP)

Berikut adalah rekam peta perjalanan progres pengerjaan fitur-fitur **Mini Bank** dalam lintasan jangka pendek hingga menengah.

## Tahap 1: Fondasi Arsitektur & Core Engine
- [x] Rancang struktur direktori standardisasi **Clean Architecture** Go.
- [x] Siapkan Makefile, Dockerfile multi-stage, dan integrasi kontainer docker-compose Postgres.
- [ ] Implementasikan skema migrasi tabel awal (`accounts`, `transactions`).
- [ ] Buat lapisan Repository dan koneksi database dengan optimasi Connection Pooling (pgx / database/sql).

## Tahap 2: Layanan Kas Fungsional (Core Services & Handlers)
- [ ] Eksekusi Usecase Manajemen Rekening & Cek Saldo.
- [ ] Konstruksi Usecase Transaksi Deposit & Withdraw.
- [ ] Implementasi Transaksi Transfer Dana Atomik berdaya tahan *Deadlock mitigation*.
- [ ] Susun penampung Router HTTP REST APIs & Middlewares (Logger, Error Recovery).

## Tahap 3: Keamanan, Audit & Skala Lanjut
- [ ] Tambahkan verifikasi parameter pengatur perlindungan ganda (`Idempotency Key`).
- [ ] Integrasikan proteksi autentikasi berprinsip Token (JWT / OAuth2).
- [ ] Buatkan modul pengekspor Laporan Bukti Transaksi bulanan berformat dokumen/PDF.
