package {{.EntityLower}}

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type {{.Entity}}Handler struct {
	{{.EntityLower}}Service Service
}

func New{{.Entity}}Handler({{.EntityLower}}Service Service) *{{.Entity}}Handler {
	return &{{.Entity}}Handler{ {{.EntityLower}}Service: {{.EntityLower}}Service }
}

func (h *{{.Entity}}Handler) Init(c *gin.Context) {
	// Remove this function if not needed
	c.JSON(http.StatusOK, gin.H{"status": "Hi {{.Entity}}"})
}
