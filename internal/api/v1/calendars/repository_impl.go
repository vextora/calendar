package calendars

import (
	"gorm.io/gorm"
)

type calendarsRepository struct {
	db *gorm.DB
}

func NewCalendarsRepository(db *gorm.DB) Repository {
	return &calendarsRepository{db}
}
