package auditlogger

import (
	"bytes"
	"calendarapi/internal/api/v1/auditlog"
	"calendarapi/pkg/shared"
	"encoding/json"
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type AuditLoggerMiddleware struct {
	Service auditlog.Service
	Logger  *zap.Logger
}

func NewAudtLoggerMiddleware(service auditlog.Service, logger *zap.Logger) *AuditLoggerMiddleware {
	return &AuditLoggerMiddleware{
		Service: service,
		Logger:  logger,
	}
}

func (m *AuditLoggerMiddleware) Handle() gin.HandlerFunc {
	return func(c *gin.Context) {
		/*if c.Request.Method != http.MethodPost && c.Request.Method != http.MethodPut && c.Request.Method != http.MethodDelete {
			c.Next()
			return
		}*/

		var bodyBytes []byte
		if c.Request.Body != nil {
			bodyBytes, _ = io.ReadAll(c.Request.Body)
			c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
		}

		start := time.Now()
		c.Next()
		duration := time.Since(start)

		userIDStr, ok := c.Get(shared.ContextKeyUserID)
		if !ok {
			m.Logger.Warn("user ID tidak ditemukan di context")
			return
		}

		userID, ok := userIDStr.(int)
		if !ok {
			m.Logger.Warn("user ID bukan bertipe uint", zap.Any("user_id", userID))
			return
		}

		action := strings.ToLower(c.Request.Method)
		entity := strings.Split(strings.Trim(c.Request.URL.Path, "/"), "/")
		fmt.Println(c.Request.URL.Path)
		entityName := ""
		if len(entity) > 2 {
			entityName = entity[2]
		}

		statusCode := c.Writer.Status()
		if statusCode >= 400 {
			return
		}

		traceID := ""
		if tid, ok := c.Get("trace_id"); ok {
			traceID, _ = tid.(string)
		}

		detail := map[string]interface{}{
			"request_body": string(bodyBytes),
			"duration_ms":  duration.Milliseconds(),
			"status_code":  statusCode,
			"path":         c.Request.URL.Path,
			"method":       c.Request.Method,
			"ip":           c.ClientIP(),
			"user_agent":   c.Request.UserAgent(),
			"trace_id":     traceID,
		}

		detailBytes, err := json.Marshal(detail)
		if err != nil {
			m.Logger.Warn("gagal marshal audit log detail", zap.Error(err))
			return
		}

		err = m.Service.Log(
			c.Request.Context(),
			userID,
			action,
			entityName,
			"",
			detailBytes,
		)
		if err != nil {
			m.Logger.Error("gagal log audit", zap.Error(err))
		} else {
			m.Logger.Info("audit log",
				zap.String("log_type", "audit"),
				zap.String("trace_id", traceID),
				zap.String("action", action),
				zap.Int("user_id", userID),
				zap.String("resource", entityName),
				zap.String("path", c.Request.URL.Path),
				zap.Any("payload", json.RawMessage(detailBytes)),
			)
		}
	}
}
