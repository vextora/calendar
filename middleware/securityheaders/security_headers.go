package securityheaders

import (
	"oncomapi/pkg/config"
	"oncomapi/pkg/securityheader"

	"github.com/gin-gonic/gin"
)

func SecurityHeadersMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		opts := securityheader.Options{
			XFrameOptions:           config.GetEnvString(config.ShXFrameOptions),
			XContentTypeOptions:     config.GetEnvString(config.ShXContentTypeOptions),
			XXSSProtection:          config.GetEnvString(config.ShXxssProtection),
			ContentSecurityPolicy:   config.GetEnvString(config.ShContentSecurityPolicy),
			ReferrerPolicy:          config.GetEnvString(config.ShReferrerPolicy),
			StrictTransportSecurity: config.GetEnvString(config.ShStrictTransportSecurity),
		}

		headers := map[string]string{}
		securityheader.ApplySecurityHeaders(headers, opts)
		for k, v := range headers {
			c.Writer.Header().Set(k, v)
		}
		c.Next()
	}
}
