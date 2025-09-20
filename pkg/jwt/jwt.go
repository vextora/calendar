package jwt

import (
	"calendarapi/pkg/config"
	"errors"
	"time"

	jwtlib "github.com/golang-jwt/jwt/v5"
)

var (
	secretKey    []byte
	appName      string
	tokenExpired int
	loaded       bool
)

func loadEnv() {
	if loaded {
		return
	}

	secretKey = []byte(config.GetEnvString(config.JwtSecret))
	appName = config.GetEnvString(config.AppName)
	tokenExpired = config.GetEnvInt(config.JwtTokenExpired)
	loaded = true
}

func GenerateToken(userID int) (string, error) {
	loadEnv()
	claims := CustomClaims{
		UserID: userID,
		RegisteredClaims: jwtlib.RegisteredClaims{
			ExpiresAt: jwtlib.NewNumericDate(time.Now().Add(time.Duration(tokenExpired) * time.Minute)),
			IssuedAt:  jwtlib.NewNumericDate(time.Now()),
			Issuer:    appName,
		},
	}
	//logs.Debug("waktu skrg : ", time.Now())
	//logs.Debug("waktu duration : ", time.Now().Add(time.Duration(tokenExpired)*time.Hour))
	//logs.Debug("claims : %+v\n", claims)

	token := jwtlib.NewWithClaims(jwtlib.SigningMethodHS256, claims)
	return token.SignedString(secretKey)
}

func ValidateToken(tokenStr string) (*CustomClaims, error) {
	loadEnv()
	token, err := jwtlib.ParseWithClaims(tokenStr, &CustomClaims{}, func(t *jwtlib.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil || !token.Valid {
		return nil, errors.New("token is not valid")
	}

	claims, ok := token.Claims.(*CustomClaims)
	if !ok {
		return nil, errors.New("cannot get claims")
	}

	if claims.UserID == 0 {
		return nil, errors.New("invalid user ID")
	}

	return claims, nil
}
