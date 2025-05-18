package router

import (
	"fmt"
	"oncomapi/internal/api/v1/auditlog"
	"oncomapi/internal/app"
	"oncomapi/middleware/auditlogger"
	"oncomapi/pkg/observability/metrics"
	"oncomapi/pkg/router"

	_ "oncomapi/routerloader"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func InitRouter(r *gin.Engine, db *gorm.DB, auditSvc auditlog.Service, logger *zap.Logger) *gin.Engine {
	router.SetDB(db)

	fmt.Println("Mulai daftarin auto logger router...")
	auditLogger := auditlogger.NewAudtLoggerMiddleware(auditSvc, logger)
	r.Use(auditLogger.Handle())
	fmt.Println("Selesai daftarin auto logger router...")

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong pong"})
	})

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.GET("/health", app.HealthCheck)
	r.GET("/healthz", app.HealthzCheck)
	r.GET("/readyz", app.ReadinessCheck(db))
	r.GET("/metrics", gin.WrapH(metrics.MetricsHandler()))

	api := r.Group("/api")
	router.RegisterRoutes(api)

	return r
}
