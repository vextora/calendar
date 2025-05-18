package article

import (
	"oncomapi/middleware/auth"
	"oncomapi/pkg/jwt"
	"oncomapi/pkg/router"

	"github.com/gin-gonic/gin"
)

func init() {
	router.Register(RegisterArticleRoutes)
}

func RegisterArticleRoutes(rg *gin.RouterGroup) {
	db := router.GetDB()
	handler := Init(db)
	jwtValidator := jwt.NewJWTValidator()

	group := rg.Group("/v1/articles")

	group.Use(auth.JWTMiddleware(jwtValidator))
	{
		group.GET("/", handler.GetAll)
		group.GET("/:id", handler.GetByID)
		group.POST("/", handler.Create)
		group.PUT("/:id", handler.Update)
		group.DELETE("/:id", handler.Delete)

		group.GET("/my", handler.GetUserArticles)
	}
}
