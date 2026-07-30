# Design Document (System & Interaction Flows)

## 1. Alur Transaksi Atomik (Transfer Dana)

Dalam transaksi pemindahan dana (Transfer) dari **Account A** ke **Account B**, sistem menerapkan pola penanganan **ACID Database Transaction** guna mencegah kegagalan konsistensi saldo atau fenomena *race condition*.

### Langkah Eksekusi (Sequence)
1. Inisiasi Database Transaction (`BEGIN TRANSACTION`).
2. Lock baris saldo untuk Rekening Pengirim (Account A) & Rekening Penerima (Account B) menggunakan spesifiksi `FOR UPDATE` (untuk Postgres), di urutkan berdasarkan `ID` terkecil untuk mencegah **Deadlock**.
3. Validasi bahwa kecukupan saldo (Balance >= Amount) untuk Account A dipenuhi.
4. Pengurangan saldo dari Account A (`UPDATE accounts SET balance = balance - amount WHERE id = A`).
5. Penambahan saldo ke Account B (`UPDATE accounts SET balance = balance + amount WHERE id = B`).
6. Pencatatan log rekam medis finansial (Ledger entries) untuk debit di A dan kredit di B.
7. Komit transaksi (`COMMIT`).

## 2. Penanganan Kegagalan (Error & Rollback)
Jika salah satu tahap di atas gagal (contoh: gangguan jaringan DB, constraint saldo minus), seluruh eksekusi dikembalikan ke kondisi mula-mula menggunakan `ROLLBACK` dan menghasilkan *Error Message* standar (contoh: `ERR_INSUFFICIENT_FUNDS` atau `ERR_DB_CONCURRENCY`).
