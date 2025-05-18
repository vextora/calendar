package logger

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

const ContextLoggerKey = "logger"

func ZapRequestLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		reqLogger := GetLoggerWithTraceID(c)

		reqLogger = reqLogger.With(
			zap.String("method", c.Request.Method),
			zap.String("path", c.Request.URL.Path),
		)

		if userID, ok := c.Get("userID"); ok {
			if uid, ok := userID.(string); ok {
				reqLogger = reqLogger.With(zap.String("user_id", uid))
			}
		}

		c.Set(ContextLoggerKey, reqLogger)

		c.Next()
	}
}

func FromContext(c *gin.Context) *zap.Logger {
	l, ok := c.Get(ContextLoggerKey)
	if !ok {
		return Logger
	}
	if logger, ok := l.(*zap.Logger); ok {
		return logger
	}
	return Logger
}
