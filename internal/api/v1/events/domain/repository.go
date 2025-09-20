package domain

import "time"

type Repository interface {
	GetAll() ([]*Events, error)
	GetByRange(startDate, endDate time.Time) ([]*Events, error)
}
