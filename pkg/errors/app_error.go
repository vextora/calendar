package apperror

import (
	"fmt"
	"net/http"
)

type AppError struct {
	Code    int
	Message string
	Details interface{}
}

func (e *AppError) Error() string {
	return fmt.Sprintf("Code: %d, Message: %s, Details: %v", e.Code, e.Message, e.Details)
}

func (e *AppError) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"code":    e.Code,
		"status":  http.StatusText(e.Code),
		"message": e.Message,
		"details": e.Details,
	}
}

func NotFound(resource string, id interface{}) *AppError {
	return &AppError{
		Code:    404,
		Message: fmt.Sprintf("%s not found", resource),
		Details: map[string]interface{}{"id": id},
	}
}

func Internal(err error) *AppError {
	return &AppError{
		Code:    500,
		Message: "Internal server error",
		Details: err.Error(),
	}
}
