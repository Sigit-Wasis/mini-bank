# 🏦 Mini Bank App

Mini Bank App adalah aplikasi perbankan digital sederhana yang mensimulasikan fitur core banking seperti autentikasi pengguna, manajemen saldo, transfer dana, top up saldo melalui payment gateway, dan riwayat transaksi.

Aplikasi ini dirancang dengan pendekatan **modular architecture** dan terintegrasi dengan **REST API Mini Bank**.

---

## ✨ Fitur Utama

### 🔐 1. Authentication Module (Gerbang Masuk)

Fitur autentikasi untuk mengakses sistem.

**Fitur:**

- ✅ Register akun baru (Nama, Email, Password)
- ✅ Login dengan JWT Authentication
- ✅ Penyimpanan token (LocalStorage / Cookie)
- ✅ Auto-login saat aplikasi dibuka ulang
- ✅ Validasi token otomatis

---

### 🏠 2. Dashboard (Beranda Utama)

Halaman utama yang menampilkan kondisi finansial pengguna.

**Fitur:**

- 💳 **Balance Card**
  - Menampilkan saldo terkini dari endpoint `/accounts`
  - Format mata uang Rupiah (Rp 50.000)
  - UI gradient modern (inspirasi Kitabisa & Gojek)

- ⚡ **Quick Actions**
  - Transfer Dana
  - Isi Saldo (Top Up)
  - Riwayat Transaksi

---

### 💸 3. Transfer Module (Core Banking)

Fitur pengiriman dana antar akun.

**Fitur:**

- Input nomor rekening / nomor HP tujuan
- 🔍 Validasi nama penerima secara real-time
- Input nominal dengan numeric keypad
- Tombol nominal cepat (50rb, 100rb, 200rb)
- Halaman konfirmasi transfer
- Notifikasi sukses / gagal
- Validasi saldo otomatis

---

### 💰 4. Top Up Module (Payment Gateway Integration)

Pengisian saldo melalui integrasi payment gateway.

**Fitur:**

- Pilihan nominal cepat (20rb, 50rb, 100rb, 500rb)
- Pilihan metode pembayaran:
  - Virtual Account (BCA, Mandiri, dll)
  - QRIS
  - E-Wallet (GoPay, dll)

- Halaman instruksi pembayaran
- QR Code / Nomor VA
- Timer hitung mundur pembayaran
- Auto refresh status pembayaran

---

### 📊 5. Transaction History (Ledger)

Riwayat lengkap aktivitas keuangan pengguna.

**Fitur:**

- Daftar transaksi scrollable
- Pagination dari endpoint `/mutations`
- Indikator visual:
  - 🟢 Hijau → Pemasukan
  - 🔴 Merah → Pengeluaran

- Detail transaksi:
  - Tanggal & waktu
  - Jenis transaksi
  - ID transaksi
  - Saldo akhir

---

### 👤 6. Profile & Account Settings

Pengelolaan informasi akun pengguna.

**Fitur:**

- Menampilkan profil user
- Data: Nama Lengkap, Email, Username
- Ganti password (opsional)
- Logout (hapus token & redirect ke login)

---

## 🧩 Arsitektur Modul

```
modules/
│
├── auth
├── dashboard
├── transfers
├── topup
├── transactions
└── profile
```

---

## 🔗 API Endpoints

| Module        | Endpoint                        |
| ------------- | ------------------------------- |
| Auth          | `/auth/login`, `/auth/register` |
| Accounts      | `/accounts`                     |
| Transfer      | `/transfers`                    |
| Check Account | `/accounts/check`               |
| Top Up        | `/topup`                        |
| Mutations     | `/mutations`                    |
| User Profile  | `/users/me`                     |

---

## 🛠 Tech Stack

**Frontend**

- React
- React Router
- Axios
- Tailwind / CSS Framework

**Backend**

- Go Fiber
- JWT Authentication
- PostgreSQL

**Payment Integration**

- Midtrans / Payment Gateway

---

## 🎯 Tujuan Project

Project ini dibuat untuk:

- simulasi sistem perbankan digital
- pembelajaran arsitektur backend & frontend modular
- demonstrasi integrasi payment gateway
- portofolio backend & fullstack engineering

---

## 🚀 Future Improvements

- Notifikasi real-time (WebSocket)
- Biometric login
- Dark mode UI
- Multi-account support
- Role-based access control
- Audit logging

---

## 📄 License

MIT License
