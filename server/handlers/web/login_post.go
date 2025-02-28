package web

import (
	"log/slog"
	"net/http"
)

type LoginFormHandler struct{}

func NewLoginFormHandler() *LoginFormHandler {
	return &LoginFormHandler{}
}

func (h *LoginFormHandler) ServeHTTP(w http.ResponseWriter, r *http.Response) {
	err := r.Request.ParseForm()
	if err != nil {
		slog.Error(err.Error())
		http.Error(w, "Form error", http.StatusBadRequest)
		return
	}

	email := r.Request.FormValue("email")
	password := r.Request.FormValue("password")

	println("email: " + email)
	println("password: " + password)
	w.WriteHeader(http.StatusOK)
}
