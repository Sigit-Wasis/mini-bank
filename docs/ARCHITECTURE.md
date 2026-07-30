# Arsitektur Sistem (Mini Bank App Architecture)

Proyek **Mini Bank** menggunakan pendekatan **Modular Architecture** di ekosistem **Golang** yang didesain untuk siap dideploy ke infrastruktur berbasis **Docker & Kubernetes di VPS**. Arsitektur ini mengedepankan pemisahan tanggung jawab (*separation of concerns*), skalabilitas horizontal, dan observabilitas yang tinggi.

## 1. Spesifikasi Tech Stack
- **Language:** Golang (Go 1.24+)
- **HTTP Web Framework / Router:** Go Fiber (Fast & lightweight REST API router)
- **Data Access Layer / Repository:** SQLC + PGX (Type-safe SQL & performant PostgreSQL driver)
- **Primary Database:** PostgreSQL (Untuk penyimpanan relasional permanen & ACID transaction)
- **Caching & Session Storage:** Redis (Untuk Session Store, Cache, Rate Limiter, & Temporary Data)
- **Shared Modules / Utility:** Viper (Config Parsing), Zap / Slog (Structured Logging)
- **Background Worker:** Custom Go Worker / Queue processing (Async email/SMS, Webhook callbacks, Retry failed payments)
- **Edge / Ingress:** Nginx Ingress (Reverse Proxy + SSL via Let's Encrypt / Cert-Manager)
- **Monitoring & Observability:** Prometheus (Metrics), Grafana (Dashboards), Loki (Log Aggregation)

---

## 2. Struktur Direktori (Modular Architecture)

```text
mini-bank/
├── cmd/
│   ├── api/           # EntryPoint untuk server HTTP REST API (Go Fiber Router)
│   └── worker/        # EntryPoint untuk Background Worker (Async job & webhook engine)
├── configs/           # Konfigurasi file (Viper YAML/ENV definitions)
├── database/          # File migrasi PostgreSQL (up/down SQL) & query sqlc
├── deployment/        # Kubernetes manifests (StatefulSet, Deployment, Service, Ingress, CI/CD)
├── docs/              # Dokumentasi arsitektur, PRD, spesifikasi API, & panduan testing
├── internal/          # Domain logika bisnis & Modular Services
│   ├── auth/          # Authentication & Token Service (JWT, Refresh Token)
│   ├── customer/      # Customer Profiling Service
│   ├── account/       # Account Management & Saldo Service
│   ├── transfer/      # Internal Transfer & Inter-Account Service (ACID Tx)
│   ├── transaction/   # Transaction History & Statement Service
│   ├── payment/       # Payment Order & Settlement Service
│   ├── topup/         # Payment Gateway Top Up Service (Midtrans / Xendit)
│   ├── notification/  # Notification trigger (Email / SMS / Push)
│   ├── admin/         # Admin Dashboard & Monitoring Service
│   └── reporting/     # Financial Reporting & Export Service
├── pkg/               # Shared modules / pustaka lintas service
│   ├── config/        # Konfigurasi parser dengan Viper
│   ├── logger/        # Structured logging dengan Zap / Slog
│   ├── middleware/    # Go Fiber Middlewares (Auth, Rate Limiter, CORS, Error Handler)
│   └── util/          # Helper (Cryptography, Decimal formatting, ID generator)
└── scripts/           # Script pendukung automasi development
```

---

## 3. Komponen Sistem & Integrasi

### A. Client Apps & Edge Ingress
- **Mobile App (Flutter)** dan **Web Dashboard (Admin)** berkomunikasi via HTTPS secara eksklusif ke **Nginx Ingress** yang berfungsi sebagai Reverse Proxy dan pemegang sertifikat SSL (Let's Encrypt / Cert-Manager).
- Permintaan kemudian diteruskan ke **API Gateway / Router** di dalam cluster.

### B. Backend Services (Go Fiber + SQLC)
- Setiap permintaan HTTP divalidasi oleh **Middleware** (CORS, Rate Limit melalui Redis, dan Auth JWT Bearer).
- Logic dieksekusi di dalam masing-masing Service Module (`auth`, `account`, `transfer`, dll.) menggunakan prinsip isolasi domain.
- Komunikasi database menggunakan pola **Repository Pattern** yang di-generate oleh **SQLC** dengan koneksi driver **PGX** ber-pool tinggi.

### C. Background Worker & Asynchronous Tasks
- Pekerjaan yang membutuhkan durasi panjang atau komunikasi pihak ketiga diserahkan ke **Background Worker** melalui Redis Queue:
  - Mengirim notifikasi (Email / SMS Gateway).
  - Memproses webhook balasan dari **Payment Gateway (Midtrans / Xendit)**.
  - Melakukan *retry* atas order pembayaran yang tertunda/gagal.
  - Menjalankan cron *Scheduled Job* & Rekap generate laporan harian.

### D. Observability Layer
- Server API & Worker mengarahkan log berformat JSON terstruktur (*Zap/Slog*) ke **Loki** dan mengumpulkan metrik kinerja perbankan ke **Prometheus** untuk difilter dan divisualisasikan melalui **Grafana**.
