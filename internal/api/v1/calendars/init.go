package calendars

import "gorm.io/gorm"

func Init(db *gorm.DB) *CalendarsHandler {
	repo := NewCalendarsRepository(db)
	service := NewCalendarsService(repo)

	return NewCalendarsHandler(service)
}
