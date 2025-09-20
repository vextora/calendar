package observability

import (
	"calendarapi/pkg/config"
	"calendarapi/pkg/shared"

	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
)

func TraceMiddleware() gin.HandlerFunc {
	serviceName := config.GetEnvString(config.TracerServiceName)
	tracer := otel.Tracer(serviceName)

	return func(c *gin.Context) {
		ctx := c.Request.Context()
		ctx, span := tracer.Start(ctx, c.Request.Method+" "+c.FullPath())
		defer span.End()

		span.SetAttributes(
			attribute.String("http.method", c.Request.Method),
			attribute.String("http.url", c.Request.URL.Path),
			attribute.String("http.user_agent", c.Request.UserAgent()),
			attribute.String("http.client_ip", c.ClientIP()),
		)

		traceID := span.SpanContext().TraceID().String()
		c.Set(shared.TraceIDKey, traceID)

		c.Request = c.Request.WithContext(ctx)

		c.Next()
	}
}
