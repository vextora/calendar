package dto

type EventsResponse struct {
	ID uint `json:"id"`
}

type GroupedEvent struct {
	Date   string                  `json:"date"`
	Events []CalendarEventResponse `json:"events"`
}

type CalendarEventResponse struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Start       string `json:"start"`
	End         string `json:"end"`
	AllDay      bool   `json:"all_day"`
	Location    string `json:"location,omitempty"`
}
