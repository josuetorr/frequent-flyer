package utils

import "golang.org/x/crypto/bcrypt"

func ComparePassword(hashed string, psswrd string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashed), []byte(psswrd))
}
