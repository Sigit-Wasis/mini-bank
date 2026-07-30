# Aturan Pengembangan (Rules & Guidelines)

## 1. Standar Penulisan Kode (Go Coding Standards)
- **Formatting:** Gunakan selalu pembatasan standar `gofmt` dan `goimports` untuk merapikan sintaks go.
- **Error Handling:** Jangan pernah menyembunyikan atau mengabaikan *errors*. Tulis penjelas konteks yang lugas (contoh: `fmt.Errorf("failed to deposit balance: %w", err)`).
- **Concurrency & Transaction Logging:** Pastikan setiap panggilan fungsionalitas mutasi kas dibalut didalam `context.Context` beriringan timeout serta `sql.Tx` atau koneksi yang di-passing dengan benar.

## 2. Larangan Domain Bisnis Perbankan (Financial Rules)
- **Floating Point Numbers:** DILARANG KERAS memvalidasi saldo tunai (currency currency computation) menggunakan `float32` / `float64` bawaan primitif yang berpotensi melahirkan *rounding issues*. Gunakan tipe data desimal presisi tak terbatas (misal: `github.com/shopspring/decimal`, integer dalam sen/satuan terkecil tunai, atau representasi *big int*).
- **Hard Delete Rekening:** Jangan menghapus data rekening finansial yang sudah memiliki log riwayat mutasi dari database (lakukan Soft Delete atau penguncian parameter status).
- **Log Data Sensitif:** Hindari mengekstrak kata sandi sandi, token sandi autentikasi, ataupun PIN nasabah di sistem logging (stdout/syslog/stderr).
