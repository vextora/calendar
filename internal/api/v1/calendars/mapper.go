package calendars

import "calendarapi/internal/api/v1/calendars/dto"

func CalendarsResponse(data *Calendars) dto.CalendarsResponse {
	return dto.CalendarsResponse{
		ID:      data.ID,
	}
}
