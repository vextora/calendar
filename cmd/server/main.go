// @title CalendarAPI
// @version 1.0
// @description This is the API documentation for CalendarAPI.
// @host localhost:8080
// @BasePath /api
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.

package main

import (
	"calendarapi/internal/api/v1/auditlog"
	appmigrate "calendarapi/internal/app/migration"
	appseeder "calendarapi/internal/app/seeder"
	"calendarapi/middleware/cors"

	"calendarapi/middleware/securityheaders"
	"calendarapi/pkg/config"
	"calendarapi/pkg/db"
	"calendarapi/pkg/logger"
	logs "calendarapi/pkg/logutil"

	"calendarapi/pkg/validation"
	"calendarapi/router"
	"os"
	"time"

	_ "calendarapi/docs"
	_ "calendarapi/pkg/shared"

	"github.com/gin-gonic/gin"
)

func main() {
	// Set timezone
	loc, _ := time.LoadLocation("Asia/Jakarta")
	time.Local = loc

	// Initialize configuration dan logging
	config.InitDotenv()
	logger.InitZap()

	// Database and seeder
	initDb, _ := db.InitDB()
	appmigrate.MigratePostgres(initDb.Postgres)
	appseeder.SeedAll(initDb.Postgres)

	// Validator and Rate limiter
	validation.InitValidator()

	// Initialize Gin and middleware
	r := gin.New()
	r.Use(cors.CorsMiddleware())                       // CORS
	r.Use(securityheaders.SecurityHeadersMiddleware()) // Security headers
	r.Use(logger.ZapRequestLogger())                   // Start log request
	r.Use(logger.ZapAccessLogger())                    // End log request (duration, status, etc)

	auditHandler := auditlog.Init(initDb.Postgres)
	auditSvc := auditHandler.Service
	zapLogger := logger.Logger.Named("audit-middleware")

	// Routing
	router.InitRouter(r, initDb.Postgres, auditSvc, zapLogger)

	// Start server
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}

	err := r.Run(":" + port)
	if err != nil {
		logs.Error("failed to run server: %v", err)
	}
}
