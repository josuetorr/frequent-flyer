package utils

import "os"

func GetJWTSecret() string {
	return os.Getenv("JWT_SECRET")
}

func GetAppName() string {
	return os.Getenv("APP_NAME")
}
