# Package (pkg)

## Deskripsi
Folder `pkg` berisi library atau utilitas yang dapat digunakan ulang di seluruh proyek.  
Biasanya berisi helper, middleware generik, logger, konfigurasi, validator, dll.

## Contoh Isi
- Logger (logging utilitas)  
- Config (parsing konfigurasi)  
- Validator (validasi input)  
- Middleware generik (jika tidak di folder middleware khusus)

## Cara Pakai
Import package dari `pkg` kapanpun dibutuhkan untuk menghindari duplikasi kode.

## Catatan
- Usahakan package di `pkg` tidak bergantung pada domain spesifik di `internal`.
- Buat package yang reusable dan generic.
