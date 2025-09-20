package events

import (
	"time"

	"gorm.io/gorm"
)

type eventsRepository struct {
	db *gorm.DB
}

func NewEventsRepository(db *gorm.DB) Repository {
	return &eventsRepository{db}
}

func (r *eventsRepository) GetAll() ([]*Events, error) {
	var events []*Events
	err := r.db.Find(&events).Error
	return events, err
}

func (r *eventsRepository) GetByRange(startDate, endDate time.Time) ([]*Events, error) {
	var events []*Events
	err := r.db.Where("start_time <= ? AND end_time >= ?", endDate, startDate).Find(&events).Error
	if err != nil {
		return nil, err
	}

	return events, nil
}
