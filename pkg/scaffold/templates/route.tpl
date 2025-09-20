package {{.EntityLower}}

import (
	// "calendarapi/middleware/auth"
	// "calendarapi/pkg/jwt"
	"calendarapi/pkg/router"

	"github.com/gin-gonic/gin"
)

func init() {
	router.Register(Register{{.Entity}}Routes)
}

func Register{{.Entity}}Routes(rg *gin.RouterGroup) {
	db := router.GetDB()
	handler := Init(db)
	// jwtValidator := jwt.NewJWTValidator()
	group := rg.Group("/{{.Version}}/{{.EntityLower}}")

	group.GET("/", handler.Init)

	// group.Use(auth.JWTMiddleware(jwtValidator))
	// {
	// }
}
