package forms

import (
	"log/slog"
	"net/http"
	"net/mail"

	"github.com/josuetorr/frequent-flyer/server/handlers"
)

type LoginPostHandler struct {
	authService handlers.AuthService
}

func NewLoginPostHandler(authService handlers.AuthService) *LoginPostHandler {
	return &LoginPostHandler{authService: authService}
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

	_, err = mail.ParseAddress(email)
	if err != nil {
		slog.Error(err.Error())
		http.Error(w, "Invalid email", http.StatusBadRequest)
		return
	}

	sessionID, err := h.authService.Login(r.Context(), email, password)
	if err != nil {
		slog.Error(err.Error())
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	println("session id: " + sessionID)

	w.WriteHeader(http.StatusOK)
}
