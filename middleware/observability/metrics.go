package observability

import (
	"calendarapi/pkg/observability/metrics"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func PrometheusMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()

		duration := time.Since(start).Seconds()
		status := strconv.Itoa(c.Writer.Status())
		method := c.Request.Method
		endpoint := c.FullPath()
		if endpoint == "" {
			endpoint = c.Request.URL.Path
		}

		metrics.RequestCount.WithLabelValues(method, endpoint, status).Inc()
		metrics.RequestDuration.WithLabelValues(method, endpoint, status).Observe(duration)
	}
}
