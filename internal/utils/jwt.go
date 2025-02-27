package utils

import (
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/josuetorr/frequent-flyer/internal/data"
)

func NewAccessToken(userId data.ID) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss": GetAppName(),
		"sub": userId,
		"exp": time.Now().Add(time.Minute * 15).Unix(),
	})
	return token.SignedString([]byte(GetJwtAccessSecret()))
}

func NewRefreshToken(userId data.ID) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss": GetAppName(),
		"sub": userId,
		"exp": time.Now().Add(time.Hour * 24 * 7).Unix(),
	})
	return token.SignedString([]byte(GetJwtRefreshSecret()))
}
