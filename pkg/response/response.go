package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func JsonResponse(code int, status string, data interface{}, message interface{}) map[string]interface{} {
	return map[string]interface{}{
		"code":    code,
		"status":  status,
		"data":    data,
		"message": message,
	}
}

func Success(data interface{}, message string) map[string]interface{} {
	return JsonResponse(http.StatusOK, http.StatusText(http.StatusOK), data, message)
}

func Error(code int, message interface{}) map[string]interface{} {
	return JsonResponse(code, http.StatusText(code), struct{}{}, message)
}

func SingleMessage(message string) map[string]interface{} {
	return JsonResponse(http.StatusOK, http.StatusText(http.StatusOK), struct{}{}, map[string]string{"message": message})
}

func ValidationError(message map[string]string) map[string]interface{} {
	return JsonResponse(http.StatusBadRequest, "Bad Request", struct{}{}, message)
}

func SendError(c *gin.Context, code int, msg interface{}) {
	if msgStr, ok := msg.(string); ok {
		c.JSON(code, Error(code, msgStr))
	} else {
		c.JSON(code, Error(code, "Internal server error"))
	}
}

func SendSuccess(c *gin.Context, data interface{}, message ...string) {
	msg := ""
	if len(message) > 0 {
		msg = message[0]
	}
	c.JSON(http.StatusOK, Success(data, msg))
}

func SendMessage(c *gin.Context, message string) {
	c.JSON(http.StatusOK, SingleMessage(message))
}

func SendValidationError(c *gin.Context, messages map[string]string) {
	c.JSON(http.StatusBadRequest, ValidationError(messages))
}
