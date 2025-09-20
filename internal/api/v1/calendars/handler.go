package calendars

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type CalendarsHandler struct {
	calendarsService Service
}

func NewCalendarsHandler(calendarsService Service) *CalendarsHandler {
	return &CalendarsHandler{ calendarsService: calendarsService }
}

func (h *CalendarsHandler) Init(c *gin.Context) {
	// Remove this function if not needed
	c.JSON(http.StatusOK, gin.H{"status": "Hi Calendars"})
}
