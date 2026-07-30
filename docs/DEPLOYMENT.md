# Panduan Deployment (DEPLOYMENT)

Dokumentasi mengenai konfigurasi penyelarasan dan rilis aplikasi **Mini Bank** baik lingkungan pengembangan (Development / Local), Staging, dan Production.

## 1. Menjalankan di Local (Lewat Docker Compose)
Pastikan system terminal Anda memiliki instalasi `docker` dan `docker-compose`.

```bash
# Salin konfigurasi environment variabel
cp .env.example .env

# Jalankan orchestration server beserta database PostgreSQL
docker-compose up -d --build

# Untuk melihat riwayat log kontainer running
docker-compose logs -f app
```

## 2. Proses Build & Deployment di Server (Production Minimal setup)
- Gunakan automasi alur CI/CD dari folder `deployment/` (seperti GitHub Actions) atau deploy via skrip rsync ke target VPS.
- Aplikasi binari yang dikompilasi (menggunakan multi-stage Docker build dari `Dockerfile` root) melingkupi OS base `alpine` super ringan (~15-30MB) untuk akselerasi instan startup service time (Cold boot).
