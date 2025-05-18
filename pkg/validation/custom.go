package validation

import (
	"regexp"
	"strconv"
	"strings"
	"unicode/utf8"

	"github.com/go-playground/validator/v10"
)

func RegisterCustomRules() {
	//validate.RegisterValidation("title", ValidateTitle)
	//validate.RegisterValidation("slug", ValidateSlug)
	validate.RegisterValidation("minLength", ValidateMinLength)
	validate.RegisterValidation("minWord", ValidateMinWord)
	validate.RegisterValidation("email", ValidateEmail)
}

func ValidateTitle(fl validator.FieldLevel) bool {
	return len(fl.Field().String()) >= 3
}

func ValidateSlug(fl validator.FieldLevel) bool {
	re := regexp.MustCompile(`^a-z0-9-]+$`)
	return re.MatchString(fl.Field().String())
}

func ValidateMinLength(fl validator.FieldLevel) bool {
	param := fl.Param()
	minLength, err := strconv.Atoi(param)
	if err != nil {
		return false
	}

	if utf8.RuneCountInString(fl.Field().String()) < minLength {
		return false
	}

	return true
}

func ValidateMinWord(fl validator.FieldLevel) bool {
	param := fl.Param()
	minWords, err := strconv.Atoi(param)
	if err != nil {
		return false
	}

	wordCount := len(strings.Fields(fl.Field().String()))
	return wordCount >= minWords
}

func ValidateEmail(fl validator.FieldLevel) bool {
	email := fl.Field().String()
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
	return re.MatchString(email)
}
