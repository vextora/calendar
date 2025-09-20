package events

import "calendarapi/internal/api/v1/events/dto"

func EventsResponse(data *Events) dto.EventsResponse {
	return dto.EventsResponse{
		ID:      data.ID,
	}
}
