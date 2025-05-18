# Internal

## Deskripsi
Folder `internal` berisi implementasi domain utama aplikasi.  
Setiap subfolder di dalamnya merepresentasikan domain atau modul terpisah dengan kode bisnis (business logic), handler, repository, service, dan domain model.

## Struktur Umum
- `internal/<domain>/`  
  Contoh: `internal/article`, `internal/user`, dll  
  Mengandung domain spesifik dengan substruktur domain-driven design.

## Cara Pakai
Kode di folder `internal` biasanya tidak diekspor sebagai library, hanya untuk internal aplikasi.  
Import langsung subfolder domain sesuai kebutuhan di service atau handler lain.

## Catatan
- Pastikan menjaga batas domain agar tidak terlalu saling tergantung.  
- Semua logika bisnis utama harus ada di sini.
