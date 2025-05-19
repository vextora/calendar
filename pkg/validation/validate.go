package validation

import (
	"net/http"

	"oncomapi/pkg/response"

	"log"

	"github.com/gin-gonic/gin"
	validator "github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func InitValidator() {
	validate = validator.New()
	RegisterCustomRules()
}

func Validate(c *gin.Context, obj interface{}) bool {
	if validate == nil {
		InitValidator()
	}
	// Bind JSON request to object
	if err := c.ShouldBindJSON(obj); err != nil {
		response.SendError(c, http.StatusBadRequest, "Invalid input")
		return false
	}

	// Validate object
	if err := validate.Struct(obj); err != nil {
		errorMessages := make(map[string]string)
		for _, err := range err.(validator.ValidationErrors) {
			log.Printf("Validation error: Field=%s, Error=%s", err.Field(), err.Tag())
			msg := GetErrorMessage(err.Field(), err.Tag(), err.Param())
			if msg == "" {
				msg = err.Error()
			}

			errorMessages[err.Field()] = msg
		}
		response.SendValidationError(c, errorMessages)
		return false
	}

	return true
}
