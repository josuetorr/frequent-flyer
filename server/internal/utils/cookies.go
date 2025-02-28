package utils

import "net/http"

const SessionCookieName = "session_cookie"

func SetCookie(w http.ResponseWriter, value string, path string, age int) {
	http.SetCookie(w, &http.Cookie{
		Name:  SessionCookieName,
		Value: value,
		// HttpOnly: true,
		// Secure:   true,
		Path:   path,
		MaxAge: age,
	})
}

func InvalidateCookie(w http.ResponseWriter) {
	http.SetCookie(w, &http.Cookie{
		Name:  SessionCookieName,
		Value: "",
		// HttpOnly: true,
		// Secure:   true,
		Path:   "/",
		MaxAge: -1,
	})
}
