# Project Calendar API

## Deskripsi

Proyek ini adalah calendar API yang menggunakan recurrence_rule yang dibangun dengan bahasa Go menggunakan arsitektur domain-driven design dan modular.

## Struktur Folder

- `internal/` : Domain utama aplikasi
- `pkg/` : Utility dan package reusable
- `middleware/` : HTTP middleware
- `cmd/` : Entry point aplikasi
- `configs/` : Konfigurasi aplikasi
- `tests/` : Unit test, integration, dan e2e test

## Cara Menjalankan

1. Setup environment dan konfigurasi database
2. Jalankan perintah build dan run

```bash
make dev
```

3. Pastikan proses migration database berhasil
4. Buka postman dan gunakan http://localhost:8080/api/v1/events/all
5. Gunakan request

```
{
    "start_date": "2025-09-01",
    "end_date": "2025-09-30"
}
```

dan output seperti

```
{
    "code": 200,
    "data": [
        {
            "date": "2025-09-08",
            "events": [
                {
                    "id": 9,
                    "title": "Board Meeting",
                    "description": "Rapat dewan direksi",
                    "start": "2025-09-08T03:00:00Z",
                    "end": "2025-09-08T05:00:00Z",
                    "all_day": false,
                    "location": "Conference Hall"
                }
            ]
        },
        {
            "date": "2025-09-15",
            "events": [
                {
                    "id": 3,
                    "title": "Payroll Processing",
                    "description": "Proses gaji bulanan",
                    "start": "2025-09-15T01:00:00Z",
                    "end": "2025-09-15T03:00:00Z",
                    "all_day": false,
                    "location": "Finance Office"
                }
            ]
        },
        {
            "date": "2025-09-20",
            "events": [
                {
                    "id": 8,
                    "title": "Meditation",
                    "description": "Sesi meditasi sebelum tidur",
                    "start": "2025-09-20T22:00:00+07:00",
                    "end": "2025-09-20T22:15:00+07:00",
                    "all_day": false
                }
            ]
        },
        {
            "date": "2025-09-21",
            "events": [
                {
                    "id": 6,
                    "title": "Gym Workout",
                    "description": "Latihan rutin",
                    "start": "2025-09-21T07:00:00+07:00",
                    "end": "2025-09-21T08:00:00+07:00",
                    "all_day": false,
                    "location": "Fitness Center"
                }
            ]
        },
        {
            "date": "2025-09-22",
            "events": [
                {
                    "id": 1,
                    "title": "Daily Standup",
                    "description": "Meeting singkat harian",
                    "start": "2025-09-22T09:00:00+07:00",
                    "end": "2025-09-22T10:00:00+07:00",
                    "all_day": false,
                    "location": "Zoom"
                },
                {
                    "id": 2,
                    "title": "Team Sync",
                    "description": "Koordinasi tim",
                    "start": "2025-09-22T14:00:00+07:00",
                    "end": "2025-09-22T15:00:00+07:00",
                    "all_day": false,
                    "location": "Meeting Room 1"
                },
                {
                    "id": 4,
                    "title": "Overnight Maintenance",
                    "description": "Maintenance server",
                    "start": "2025-09-22T23:00:00+07:00",
                    "end": "2025-09-22T23:59:59+07:00",
                    "all_day": false,
                    "location": "Data Center"
                },
                {
                    "id": 10,
                    "title": "Server Maintenance",
                    "description": "Maintenance harian malam",
                    "start": "2025-09-22T23:00:00+07:00",
                    "end": "2025-09-22T23:59:59+07:00",
                    "all_day": false,
                    "location": "Data Center"
                }
            ]
        },
        {
            "date": "2025-09-23",
            "events": [
                {
                    "id": 4,
                    "title": "Overnight Maintenance",
                    "description": "Maintenance server",
                    "start": "2025-09-23T00:00:00+07:00",
                    "end": "2025-09-23T01:00:00+07:00",
                    "all_day": false,
                    "location": "Data Center"
                },
                {
                    "id": 10,
                    "title": "Server Maintenance",
                    "description": "Maintenance harian malam",
                    "start": "2025-09-23T00:00:00+07:00",
                    "end": "2025-09-23T01:00:00+07:00",
                    "all_day": false,
                    "location": "Data Center"
                }
            ]
        },
        {
            "date": "2025-09-26",
            "events": [
                {
                    "id": 7,
                    "title": "Family Dinner",
                    "description": "Makan malam keluarga",
                    "start": "2025-09-26T12:00:00Z",
                    "end": "2025-09-26T14:00:00Z",
                    "all_day": false,
                    "location": "Rumah Ibu"
                }
            ]
        },
        {
            "date": "2025-09-27",
            "events": [
                {
                    "id": 5,
                    "title": "Public Holiday",
                    "description": "Libur Nasional",
                    "start": "2025-09-27T00:00:00+07:00",
                    "end": "2025-09-27T23:59:59+07:00",
                    "all_day": true
                }
            ]
        }
    ],
    "message": "",
    "status": "OK"
}
```

## Kontak

Untuk pertanyaan lebih lanjut, hubungi tim pengembang.
