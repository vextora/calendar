package logger

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func ZapAccessLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		c.Next()

		duration := time.Since(start)
		status := c.Writer.Status()
		userID := "-"

		if l, ok := c.Get(ContextLoggerKey); ok {
			if log, ok := l.(*zap.Logger); ok {
				if tid := log.Core().Enabled(zap.InfoLevel); tid {
					if id, exists := c.Get("userID"); exists {
						userID = fmt.Sprintf("%d", id)
					}
					log.Info("Request handled",
						zap.String("log_type", "access"),
						zap.Int("status", status),
						zap.Duration("latency", duration),
						zap.String("user_id", userID),
					)
				}
			}
		}
	}
}
