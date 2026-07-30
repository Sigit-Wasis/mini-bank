# Spesifikasi API (REST API)

Seluruh endpoints bertransaksi mengggabungkan protokol **JSON over HTTPS/HTTP** dan memberikan respons berformat konsisten.

## 1. Standar Format Respons (Wrapper JSON)
```json
{
  "status": "success",
  "data": { ... },
  "error": null
}
```

## 2. Endpoints Utama

### A. Create Account
- **Method:** `POST /api/v1/accounts`
- **Request Body:**
  ```json
  {
    "owner_name": "Sigit Wasis Subekti",
    "initial_balance": 100000.00,
    "currency": "IDR"
  }
  ```
- **Response:** `201 Created`
  ```json
  {
    "status": "success",
    "data": {
      "id": "e67d2b45-6671-4d37-8822-1d54ff83c1aa",
      "account_number": "100200300400",
      "owner_name": "Sigit Wasis Subekti",
      "balance": "100000.00",
      "currency": "IDR"
    }
  }
  ```

### B. Transfer Balance (Inter-Account Transfer)
- **Method:** `POST /api/v1/transfers`
- **Request Body:**
  ```json
  {
    "from_account_id": "e67d2b45-6671-4d37-8822-1d54ff83c1aa",
    "to_account_id": "77bc54aa-2144-482d-83b3-abcfdd8e5476",
    "amount": "25000.00",
    "description": "Pembayaran jasa konsultasi"
  }
  ```
- **Response:** `200 OK`
