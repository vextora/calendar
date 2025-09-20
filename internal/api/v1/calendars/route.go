package calendars

import (
	// "calendarapi/middleware/auth"
	// "calendarapi/pkg/jwt"
	"calendarapi/pkg/router"

	"github.com/gin-gonic/gin"
)

func init() {
	router.Register(RegisterCalendarsRoutes)
}

func RegisterCalendarsRoutes(rg *gin.RouterGroup) {
	db := router.GetDB()
	handler := Init(db)
	// jwtValidator := jwt.NewJWTValidator()
	group := rg.Group("/v1/calendars")

	group.GET("/", handler.Init)

	// group.Use(auth.JWTMiddleware(jwtValidator))
	// {
	// }
}
