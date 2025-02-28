package web

import (
	"log/slog"
	"net/http"
)

type LoginPostHandler struct{}

func NewLoginPostHandler() *LoginPostHandler {
	return &LoginPostHandler{}
}

func (h *LoginPostHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		slog.Error(err.Error())
		http.Error(w, "Form error", http.StatusBadRequest)
		return
	}

	email := r.FormValue("email")
	password := r.FormValue("password")

	println("email: " + email)
	println("password: " + password)
	w.WriteHeader(http.StatusOK)
}
