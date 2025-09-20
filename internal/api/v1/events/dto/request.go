package dto

type EventsRequest struct {
	ID uint `json:"id" validate:"required"`
}

type EventRangeRequest struct {
	StartDate string `json:"start_date" validate:"required"`
	EndDate   string `json:"end_date" validate:"required"`
}
