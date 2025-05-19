package jwt

import (
	jwtlib "github.com/golang-jwt/jwt/v5"
)

type CustomClaims struct {
	UserID int `json:"user_id"`
	jwtlib.RegisteredClaims
}
