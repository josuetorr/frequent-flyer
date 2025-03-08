package utils

import "golang.org/x/crypto/bcrypt"

func ComparePassword(hashed string, psswrd string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashed), []byte(psswrd))
}

func HashPassword(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}
