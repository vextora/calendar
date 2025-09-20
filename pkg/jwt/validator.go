package jwt

import "calendarapi/pkg/shared"

type jwtValidator struct{}

func NewJWTValidator() shared.TokenValidator {
	return &jwtValidator{}
}

func (j *jwtValidator) Validate(tokenStr string) (any, error) {
	return ValidateToken(tokenStr)
}
