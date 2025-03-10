package utils

import (
	"github.com/gorilla/securecookie"
)

func EncodeCookie(cookieName string, cookieValue string) (string, error) {
	encoder := securecookie.New([]byte(GetSessionHashKey()), []byte(GetSessionBlockKey()))
	return encoder.Encode(cookieName, cookieValue)
}

func DecodeCookie(cookieName string, cookieValue string) (string, error) {
	encoder := securecookie.New([]byte(GetSessionHashKey()), []byte(GetSessionBlockKey()))
	var value string
	err := encoder.Decode(cookieName, cookieValue, &value)
	return value, err
}
