# Strategi Pengujian (TESTING)

Pengembangan sistem kas finansial menuntut pengujian cermat guna memastikan tidak terjadi kegagalan matematis dan pembatasan konkurensi (Concurrency issues).

## 1. Piramida Pengujian di Proyek Ini

### A. Unit Testing
- **Fokus:** Menguji satu logika tunggal pada *Service/Usecase Layer* (contoh: kalkulasi potongan ongkos kirim/transfer, penjaminan saldo minimum).
- **Alat:** Gunakan pemandu paket standar `testing` dari bahasa Go dikombinasikan dengan pustaka assertion semisal `github.com/stretchrify/testify/assert`.
- **Mocking:** Seluruh abstraksi panggilan ke basis data atau layanan eksternal di-mock via interface (contoh menggunakan `go.uber.org/mock` / `mockery`).

### B. Integration Testing & Concurrency Test
- **Fokus:** Menjalankan koneksi kongkrit terhadap *Ephemeral Database* (contoh via kontainer *Testcontainers Go* atau DB test Postgres tersembunyi).
- **Stress Test (Race Conditions):** Menyiapkan segerombolan *Goroutines* secara paralel menembaki eksekusi transaksi yang sama untuk menjabarkan daya tahan mitigasi penguncian DB Lock (`FOR UPDATE`).
- **Eksekusi perintah race detector:**
  ```bash
  go test -race ./...
  ```
