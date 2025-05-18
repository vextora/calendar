# Middleware

## Deskripsi
Folder ini berisi middleware HTTP yang digunakan di aplikasi, seperti autentikasi, logging, rate limiting, dll.

## Struktur
- Middleware biasanya berupa fungsi yang menerima dan mengembalikan handler HTTP.  
- Bisa digunakan di semua route atau domain tertentu.

## Cara Pakai
Import middleware dan pasang di router atau handler chain sesuai kebutuhan.

## Catatan
- Middleware harus ringan dan efisien.  
- Pisahkan middleware yang reusable dengan yang domain spesifik.
