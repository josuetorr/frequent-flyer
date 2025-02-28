package utils

import "os"

func GetJwtAccessSecret() string {
	return os.Getenv("JWT_ACCESS_SECRET")
}

func GetJwtRefreshSecret() string {
	return os.Getenv("JWT_REFRESH_SECRET")
}

func GetAppName() string {
	return os.Getenv("APP_NAME")
}

func GetAppEmail() string {
	return os.Getenv("EMAIL_APP_SENDER")
}

func GetAppEmailPassword() string {
	return os.Getenv("EMAIL_APP_PASSWORD")
}
