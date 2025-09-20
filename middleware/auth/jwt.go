package auth

import (
	"net/http"
	"strings"

	"calendarapi/internal/api/v1/user"
	"calendarapi/pkg/jwt"
	"calendarapi/pkg/response"
	"calendarapi/pkg/shared"

	"github.com/gin-gonic/gin"
)

var userService user.Service

func SetupUserService(us user.Service) {
	userService = us
}

func JWTMiddleware(tokenValidator shared.TokenValidator) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if !strings.HasPrefix(authHeader, "Bearer ") {
			response.SendError(c, http.StatusUnauthorized, "Bearer token not found or invalid format")
			c.Abort()
			return
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
		claims, err := tokenValidator.Validate(tokenStr)
		if err != nil {
			response.SendError(c, http.StatusUnauthorized, "Your session has expired. Please log in again.")
			c.Abort()
			return
		}

		customClaims, ok := claims.(*jwt.CustomClaims)
		if !ok {
			response.SendError(c, http.StatusUnauthorized, "Invalid token claims")
			c.Abort()
			return
		}
		c.Set("userID", customClaims.UserID)

		user, err := userService.FindByID(customClaims.UserID)
		if err != nil || user == nil {
			response.SendError(c, http.StatusUnauthorized, "User not found or inactive")
			c.Abort()
			return
		}
		c.Set("user", user)

		c.Next()
	}
}
