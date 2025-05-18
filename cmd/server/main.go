// @title OncomAPI
// @version 1.0
// @description This is the API documentation for OncomAPI.
// @host localhost:8080
// @BasePath /api
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.

package main

import (
	"context"
	"fmt"
	"log"
	"oncomapi/internal/api/v1/auditlog"
	appmigrate "oncomapi/internal/app/migration"
	appseeder "oncomapi/internal/app/seeder"
	"oncomapi/middleware/cors"
	"oncomapi/middleware/observability"
	ratelimit "oncomapi/middleware/rate_limit"
	"oncomapi/middleware/securityheaders"

	//ratelimit "oncomapi/middleware/rate_limit"

	"oncomapi/pkg/accesscontrol"
	"oncomapi/pkg/accesscontrol/casbin"
	"oncomapi/pkg/accesscontrol/migration"
	"oncomapi/pkg/config"
	"oncomapi/pkg/db"
	"oncomapi/pkg/logger"
	logs "oncomapi/pkg/logutil"
	"oncomapi/pkg/observability/metrics"
	"oncomapi/pkg/observability/tracing"
	"oncomapi/pkg/ratelimiter"
	"oncomapi/pkg/validation"
	"oncomapi/router"
	"os"
	"time"

	_ "oncomapi/docs"
	_ "oncomapi/pkg/shared"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println(">>>>> Hello from main.go")
	ctx := context.Background()

	// Set timezone
	loc, _ := time.LoadLocation("Asia/Jakarta")
	time.Local = loc

	// Initialize configuration dan logging
	config.InitDotenv()
	logger.InitZap()

	// Observability (metrics & tracing)
	metrics.RegisterMetric()

	tp, err := tracing.InitTracer(ctx)
	if err != nil {
		log.Fatalf("failed to initialize tracer: %v", err)
	}

	defer func() {
		if err := tp.Shutdown(ctx); err != nil {
			log.Fatalf("failed to shutdown tracer provider: %v", err)
		}
	}()

	// Database and seeder
	initDb, _ := db.InitDB()
	appmigrate.MigratePostgres(initDb.Postgres)
	appseeder.SeedAll(initDb.Postgres)

	// RBAC, ABAC, Casbin
	if err := migration.RunAllMigrations(initDb.Postgres); err != nil {
		panic(err)
	}
	if err := migration.RunAllSeeders(initDb.Postgres); err != nil {
		panic(err)
	}

	enforcer, err := casbin.NewEnforcer(initDb.Postgres)
	if err != nil {
		log.Fatal("failed to init casbin enforcer : ", err)
	}
	accesscontrol.InitGlobal(enforcer)

	// Validator and Rate limiter
	validation.InitValidator()

	// Rate limiter
	limiter := ratelimiter.InitDefaultLimiter()

	// Initialize Gin and middleware
	r := gin.New()
	//r.Use(gin.Logger())
	r.Use(cors.CorsMiddleware())                       // CORS
	r.Use(securityheaders.SecurityHeadersMiddleware()) // Security headers
	r.Use(observability.TraceMiddleware())             // Trace ID for log
	//r.Use(recovery.RecoveryWithZapAndSentry())                   // Panic recovery + log to Zap & Sentry
	r.Use(logger.ZapRequestLogger())                             // Start log request
	r.Use(ratelimit.NewRateLimiterHandler(limiter).Middleware()) // Rate limiting (Redis + Lua)
	r.Use(observability.PrometheusMiddleware())                  // Prometheus metrics
	r.Use(logger.ZapAccessLogger())                              // End log request (duration, status, etc)

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

	err = r.Run(":" + port)
	if err != nil {
		logs.Error("failed to run server: %v", err)
	}
}
