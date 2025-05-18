package jwt

import (
	"errors"

	"github.com/gin-gonic/gin"
)

func GetUserID(c *gin.Context) (int, error) {
	val, exists := c.Get("userID")
	if !exists {
		return 0, errors.New("user ID not found")
	}

	userID, ok := val.(int)
	if !ok {
		return 0, errors.New("user ID is not an int")
	}

	return userID, nil
}
