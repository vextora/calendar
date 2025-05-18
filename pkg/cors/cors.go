package cors

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func Middleware(opts Options) gin.HandlerFunc {
	return func(c *gin.Context) {
		origin := c.GetHeader("Origin")

		if origin != "" && (contains(opts.AllowedOrigins, "*") || contains(opts.AllowedOrigins, origin)) {
			c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
			c.Writer.Header().Set("Access-Control-Allow-Methods", strings.Join(opts.AllowedMethods, ", "))
			c.Writer.Header().Set("Access-Control-Allow-Headers", strings.Join(opts.AllowedHeaders, ", "))

			if opts.AllowCredentials {
				c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
			}
		}

		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}

func contains(slice []string, val string) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}
