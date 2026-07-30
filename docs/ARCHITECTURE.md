# Arsitektur Sistem (Architecture)

Projek **Mini Bank** mengadaptasi standar **Clean Architecture** dipantulkan ke kultur ekosistem **Golang**, memisahkan domain aturan bisnis sejati dari alat eksekusi teknis luar (seperti Database Driver, Protokol Transport HTTP, CLI).

## Struktur Direktori

```text
mini-bank/
├── cmd/           # Titik utama eksekusi (EntryPoint) aplikasi (mis: api server, worker, cli tool)
├── configs/       # Variabel konfigurasi lingkungan standar & struct parsing config
├── database/      # Berisi migration files (up/down SQL), seeder, dan kueri (sqlc/raw)
├── deployment/    # Konfigurasi platform CI/CD, Kubernetes manifests, nginx, script sistem devops
├── docs/          # Dokumentasi terlengkap aplikasi perbankan mini ini
├── internal/      # Kode privasi bisnis (Domain Model, Service/Usecase, Repository, Handler) yang di proteksi dari impor luar
├── pkg/           # Pustaka utilitas generik yang bisa di-import oleh proyek luar atau seluruh layer internal
└── scripts/       # Shell scripts atau tool eksternal pendukung operasional developer
```

## Layer Aplikasi (internal/)

1. **Domain Layer:** Entitas inti fungsionalitas (contoh struct `Account`, `Transaction`, interface `Repository`).
2. **Repository Layer:** Komunikasi langsung secara abstraktif dengan PostgreSQL untuk kueri eksekusi DB.
3. **Service (Usecase) Layer:** Orkestrasi logika finansial perbankan dan aturan pengecekan domain validasi.
4. **Delivery / Handler Layer:** Menangkap permintaan eksternal (contoh HTTP REST JSON API via Routing router) menuju eksekusi Usecase.
