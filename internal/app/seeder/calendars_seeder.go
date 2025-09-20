package seeder

import (
	"calendarapi/internal/api/v1/calendars"
	logs "calendarapi/pkg/logutil"

	"gorm.io/gorm"
)

func seedCalendars(db *gorm.DB) {
	var count int64

	calendar := []calendars.Calendars{
		{UserID: 1, Name: "Work", Color: "#1E90FF", IsDefault: true},
		{UserID: 1, Name: "Personal", Color: "#FF6347", IsDefault: false},
	}

	db.Model([]calendars.Calendars{}).Count(&count)

	if count == 0 {
		for _, data := range calendar {
			if err := db.Create(&data).Error; err != nil {
				logs.Error("Gagal seeding article : %v\n", err)
			}
		}
	}
}
