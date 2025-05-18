package user

import (
	"oncomapi/pkg/router"

	"github.com/gin-gonic/gin"
)

func init() {
	router.Register(RegisterUserRoutes)
}

func RegisterUserRoutes(rg *gin.RouterGroup) {
	db := router.GetDB()
	handler := Init(db)

	group := rg.Group("/users")
	{
		group.POST("/register", handler.Register)
		group.POST("/login", handler.Login)
	}
}
