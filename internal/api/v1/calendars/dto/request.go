package dto

type CalendarsRequest struct {
	ID   uint `json:"id" validate:"required"`
}