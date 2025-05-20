package oncom

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type OncomHandler struct {
	oncomService Service
}

func NewOncomHandler(oncomService Service) *OncomHandler {
	return &OncomHandler{oncomService: oncomService}
}

func (h *OncomHandler) Init(c *gin.Context) {
	// Remove this function if not needed
	c.JSON(http.StatusOK, gin.H{"status": "Hi Oncom"})
}
