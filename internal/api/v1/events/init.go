package events

import "gorm.io/gorm"

func Init(db *gorm.DB) *EventsHandler {
	repo := NewEventsRepository(db)
	service := NewEventsService(repo)

	return NewEventsHandler(service)
}
