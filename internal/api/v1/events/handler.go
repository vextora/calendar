package events

import (
	"calendarapi/internal/api/v1/events/dto"
	"calendarapi/pkg/response"
	"calendarapi/pkg/validation"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type EventsHandler struct {
	eventsService Service
}

func NewEventsHandler(eventsService Service) *EventsHandler {
	return &EventsHandler{eventsService: eventsService}
}

func (h *EventsHandler) GetAll(c *gin.Context) {
	input := dto.EventRangeRequest{}
	if !validation.Validate(c, &input) {
		return
	}

	startRange, err := time.Parse("2006-01-02", input.StartDate)
	if err != nil {
		response.SendError(c, http.StatusInternalServerError, err.Error())
		return
	}
	endRange, err := time.Parse("2006-01-02", input.EndDate)
	if err != nil {
		response.SendError(c, http.StatusInternalServerError, err.Error())
		return
	}
	endRange = endRange.Add(23*time.Hour + 59*time.Minute + 59*time.Second)

	events, err := h.eventsService.GetByRange(startRange, endRange)
	if err != nil {
		response.SendError(c, http.StatusInternalServerError, err.Error())
		return
	}

	//c.JSON(http.StatusOK, events)
	response.SendSuccess(c, events)
}
