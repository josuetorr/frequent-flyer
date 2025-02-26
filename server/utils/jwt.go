package utils

import (
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/josuetorr/frequent-flyer/server/data"
)

func NewJwtToken(userId data.ID) *jwt.Token {
	return jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss": GetAppName(),
		"sub": userId,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})
}

func SignToken(token *jwt.Token) (string, error) {
	return token.SignedString([]byte(GetJWTSecret()))
}
