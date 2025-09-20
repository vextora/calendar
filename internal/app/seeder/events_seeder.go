package seeder

import (
	"calendarapi/internal/api/v1/events"
	logs "calendarapi/pkg/logutil"

	"time"

	"gorm.io/gorm"
)

func seedEvents(db *gorm.DB) {
	var count int64

	layout := time.RFC3339

	parseTime := func(s string) time.Time {
		t, err := time.Parse(layout, s)
		if err != nil {
			panic(err)
		}

		return t
	}

	strPtr := func(s string) *string {
		return &s
	}

	event := []events.Events{
		{
			CalendarID:     1,
			Title:          "Daily Standup",
			Description:    "Meeting singkat harian",
			StartTime:      parseTime("2025-09-22T09:00:00+07:00"),
			EndTime:        parseTime("2025-09-22T10:00:00+07:00"),
			Location:       strPtr("Zoom"),
			IsAllDay:       false,
			RecurrenceRule: strPtr("FREQ=DAILY;INTERVAL=1;BYDAY=MO,TU,WE,TH,FR;UNTIL=2025-10-31"),
		},
		{
			CalendarID:     1,
			Title:          "Team Sync",
			Description:    "Koordinasi tim",
			StartTime:      parseTime("2025-09-22T14:00:00+07:00"),
			EndTime:        parseTime("2025-09-22T15:00:00+07:00"),
			Location:       strPtr("Meeting Room 1"),
			IsAllDay:       false,
			RecurrenceRule: strPtr("FREQ=WEEKLY;BYDAY=MO,FR;UNTIL=2025-12-31"),
		},
		{
			CalendarID:     1,
			Title:          "Payroll Processing",
			Description:    "Proses gaji bulanan",
			StartTime:      parseTime("2025-09-15T08:00:00+07:00"),
			EndTime:        parseTime("2025-09-15T10:00:00+07:00"),
			Location:       strPtr("Finance Office"),
			IsAllDay:       false,
			RecurrenceRule: strPtr("FREQ=MONTHLY;BYMONTHDAY=15;COUNT=6"),
		},
		{
			CalendarID:     1,
			Title:          "Overnight Maintenance",
			Description:    "Maintenance server",
			StartTime:      parseTime("2025-09-22T23:00:00+07:00"),
			EndTime:        parseTime("2025-09-23T01:00:00+07:00"),
			Location:       strPtr("Data Center"),
			IsAllDay:       false,
			RecurrenceRule: strPtr("FREQ=DAILY;INTERVAL=1;UNTIL=2025-09-30"),
		},
		{
			CalendarID:     1,
			Title:          "Public Holiday",
			Description:    "Libur Nasional",
			StartTime:      parseTime("2025-09-27T00:00:00+07:00"),
			EndTime:        parseTime("2025-09-27T23:59:59+07:00"),
			Location:       nil,
			IsAllDay:       true,
			RecurrenceRule: nil,
		},
		{
			CalendarID:     2,
			Title:          "Gym Workout",
			Description:    "Latihan rutin",
			StartTime:      parseTime("2025-09-21T07:00:00+07:00"),
			EndTime:        parseTime("2025-09-21T08:00:00+07:00"),
			Location:       strPtr("Fitness Center"),
			IsAllDay:       false,
			RecurrenceRule: strPtr("FREQ=DAILY;INTERVAL=2;UNTIL=2025-10-10"),
		},
		{
			CalendarID:     2,
			Title:          "Family Dinner",
			Description:    "Makan malam keluarga",
			StartTime:      parseTime("2025-09-26T19:00:00+07:00"),
			EndTime:        parseTime("2025-09-26T21:00:00+07:00"),
			Location:       strPtr("Rumah Ibu"),
			IsAllDay:       false,
			RecurrenceRule: strPtr("FREQ=WEEKLY;BYDAY=FR;COUNT=5"),
		},
		{
			CalendarID:     2,
			Title:          "Meditation",
			Description:    "Sesi meditasi sebelum tidur",
			StartTime:      parseTime("2025-09-20T22:00:00+07:00"),
			EndTime:        parseTime("2025-09-20T22:15:00+07:00"),
			Location:       nil,
			IsAllDay:       false,
			RecurrenceRule: strPtr("FREQ=DAILY;INTERVAL=1;UNTIL=2025-09-30"),
		},
		{
			CalendarID:     1,
			Title:          "Board Meeting",
			Description:    "Rapat dewan direksi",
			StartTime:      parseTime("2025-09-08T10:00:00+07:00"),
			EndTime:        parseTime("2025-09-08T12:00:00+07:00"),
			Location:       strPtr("Conference Hall"),
			IsAllDay:       false,
			RecurrenceRule: strPtr("FREQ=MONTHLY;BYDAY=2MO;COUNT=4"),
		},
		{
			CalendarID:     1,
			Title:          "Server Maintenance",
			Description:    "Maintenance harian malam",
			StartTime:      parseTime("2025-09-22T23:00:00+07:00"),
			EndTime:        parseTime("2025-09-23T01:00:00+07:00"),
			Location:       strPtr("Data Center"),
			IsAllDay:       false,
			RecurrenceRule: strPtr("FREQ=DAILY;INTERVAL=1;UNTIL=2025-09-30"),
		},
	}

	db.Model([]events.Events{}).Count(&count)

	if count == 0 {
		for _, data := range event {
			if err := db.Create(&data).Error; err != nil {
				logs.Error("Gagal seeding article : %v\n", err)
			}
		}
	}
}
