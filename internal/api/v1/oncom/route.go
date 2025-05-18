package oncom

import (
	// "oncomapi/middleware/auth"
	// "oncomapi/pkg/jwt"
	"oncomapi/pkg/router"

	"github.com/gin-gonic/gin"
)

func init() {
	router.Register(RegisterOncomRoutes)
}

func RegisterOncomRoutes(rg *gin.RouterGroup) {
	db := router.GetDB()
	handler := Init(db)
	// jwtValidator := jwt.NewJWTValidator()
	group := rg.Group("/v1/oncom")

	group.GET("/", handler.Init)

	// group.Use(auth.JWTMiddleware(jwtValidator))
	// {
	// }
}
