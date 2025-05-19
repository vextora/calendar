package recovery

import (
	"net/http"
	"oncomapi/pkg/logger"
	"runtime/debug"
	"time"

	sentry "github.com/getsentry/sentry-go"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func RecoveryWithZapAndSentry() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		defer func() {
			if rec := recover(); rec != nil {
				stack := string(debug.Stack())

				// Get logger with trace ID if exist
				zapLogger := logger.GetLoggerWithTraceID(c)

				// Panic log
				zapLogger.Error("panic recovered",
					zap.Any("recover", rec),
					zap.String("stack", stack),
					zap.String("method", c.Request.Method),
					zap.String("path", c.FullPath()),
				)

				// Send panic to sentry
				sentry.WithScope(func(scope *sentry.Scope) {
					scope.SetTag("type", "panic")
					scope.SetExtra("method", c.Request.Method)
					scope.SetExtra("path", c.FullPath())
					scope.SetExtra("stack", stack)
					scope.SetLevel(sentry.LevelFatal)
					sentry.CaptureException(rec.(error))
				})
				sentry.Flush(2 * time.Second)

				// Send response to client
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
					"code":    500,
					"status":  "ERROR",
					"message": "Internal server error",
				})
				return
			}
		}()

		// Run request
		c.Next()

		// Log error from c.Error(err)
		for _, ginErr := range c.Errors {
			zapLogger := logger.GetLoggerWithTraceID(c)
			zapLogger.Error("handled error",
				zap.String("method", c.Request.Method),
				zap.String("path", c.FullPath()),
				zap.Error(ginErr.Err),
			)

			sentry.WithScope(func(scope *sentry.Scope) {
				scope.SetTag("type", "handler error")
				scope.SetExtra("method", c.Request.Method)
				scope.SetExtra("path", c.FullPath())
				sentry.CaptureException(ginErr.Err)
			})
		}
		sentry.Flush(2 * time.Second)

		// General access log
		duration := time.Since(start)
		zapLogger := logger.GetLoggerWithTraceID(c)
		zapLogger.Info("request finished",
			zap.String("method", c.Request.Method),
			zap.String("path", c.FullPath()),
			zap.Int("status", c.Writer.Status()),
			zap.Duration("duration", duration),
		)
	}
}
