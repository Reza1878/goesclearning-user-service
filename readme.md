# Panduan Migrasi Database untuk User Service

Dokumentasi ini menjelaskan langkah-langkah yang perlu dilakukan untuk menjalankan migrasi dan setup database PostgreSQL untuk aplikasi User Service. Panduan ini mencakup pembuatan ekstensi \`uuid-ossp\` dan pembuatan tabel \`users\` yang diperlukan dalam aplikasi.

## 1. Persiapan

Pastikan kamu sudah memiliki PostgreSQL yang terinstal di sistem atau menggunakan Docker untuk menjalankan PostgreSQL.

Jika menggunakan Docker, kamu bisa menjalankan PostgreSQL menggunakan perintah berikut:

```bash
docker run --name postgresql -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=postgres -p 5432:5432 -v /var/lib/postgresql/data -d postgres
```

Untuk melakukan koneksi ke PostgreSQL, pastikan kamu menggunakan \`host=localhost\`, \`port=5432\`, \`user=postgres\`, dan \`password=postgres\`.

## 2. Database Name / Ekstensi UUID

Untuk mendukung penggunaan tipe data UUID di PostgreSQL, pastikan ekstensi \`uuid-ossp\` sudah diaktifkan. Jalankan perintah berikut pada database PostgreSQL untuk mengaktifkan ekstensi \`uuid-ossp\`:

```sql
CREATE DATABASE user_db;
```

```sql
CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";
```

Ekstensi ini digunakan untuk menghasilkan UUID secara otomatis.

## 3. Persiapan Tabel Users

```sql
CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(100) NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    password TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```
