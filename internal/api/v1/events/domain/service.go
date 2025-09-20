package domain

import (
	"calendarapi/internal/api/v1/events/dto"
	"time"
)

type Service interface {
	GetAll() ([]*Events, error)
	GetByRange(startDate time.Time, endDate time.Time) ([]dto.GroupedEvent, error)
}
