package accesscontrol

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type MiddlewareParams struct {
	Object string
	Action string
}

func (ac *AccessChecker) RequireAccess(params MiddlewareParams) gin.HandlerFunc {
	return func(c *gin.Context) {
		user, exists := c.Get("user")
		if !exists || user == nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			return
		}

		ok, err := ac.CheckAccess(user, params.Object, params.Action)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
			return
		}

		if !ok {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "forbidden"})
			return
		}
		c.Next()
	}
}
