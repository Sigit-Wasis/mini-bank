# Security & Data Integrity

Karena Mini Bank merepresentasikan model finansial, aspek keamanan komputasi merupakan parameter kritis.

## 1. Pengamanan Komunikasi & Autentikasi
- **TLS/HTTPS Encryption:** Seluruh lalu lintas request client di-wajibkan melewati jalur terenkripsi Transport Layer Security (TLS).
- **JWT / Bearer Tokens:** Setiap otorisasi akses menggunakan tanda pengenal Token rahasia yang memiliki batas waktu usia pakai (Short-Lived Expiration Time).
- **Role-Based Access Control (RBAC):** Memisahkan otorisasi tindakan murni milik pribadi (Nasabah) berbanding tindakan kelola administratif system (Teller / Administrator).

## 2. Proteksi Serangan Eksternal
- **Idempotency Key (Double Request Shield):** Transaksi POST sensitif (Deposit & Transfer) wajib mendukung header `X-Idempotency-Key` untuk menghempas potensi gandanya pemanggilan API akibat *network retries* atau kelalaian klik dari client.
- **Rate Limiting & Throttle:** Membatasi lonjakan pemanggilan request mencurigakan (DDoS mitigations) per-IP / per-Akun per satuan waktu pendek.
- **SQL Injection Safeguards:** Penyerahan parameter kueri absolut melalui mekanisme prepared statements (contoh: `$1`, `$2` di Postgres / parameterisasi GORM/sqlc).
