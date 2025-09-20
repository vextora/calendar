package domain

import "time"

type Events struct {
	ID             uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	CalendarID     int       `json:"calendar_id"`
	Title          string    `json:"title" validate:"required"`
	Description    string    `json:"description,omitempty"`
	StartTime      time.Time `json:"start_time"`
	EndTime        time.Time `json:"end_time"`
	Location       *string   `json:"location,omitempty"`
	IsAllDay       bool      `json:"is_all_day"`
	RecurrenceRule *string   `json:"recurrence_rule,omitempty"`
	CreatedAt      time.Time `json:"created_at"`
}
