package events

import (
	// "calendarapi/middleware/auth"
	// "calendarapi/pkg/jwt"
	"calendarapi/pkg/router"

	"github.com/gin-gonic/gin"
)

func init() {
	router.Register(RegisterEventsRoutes)
}

func RegisterEventsRoutes(rg *gin.RouterGroup) {
	db := router.GetDB()
	handler := Init(db)
	// jwtValidator := jwt.NewJWTValidator()
	group := rg.Group("/v1/events")
	{
		group.GET("/all", handler.GetAll)
	}

	// group.Use(auth.JWTMiddleware(jwtValidator))
	// {
	// }
}
