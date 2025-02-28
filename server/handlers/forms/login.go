package forms

import (
	"log/slog"
	"net/http"
	"net/mail"

	"github.com/josuetorr/frequent-flyer/server/handlers"
	"github.com/josuetorr/frequent-flyer/server/internal/utils"
)

type LoginPostHandler struct {
	authService handlers.AuthService
}

func NewLoginHandler(authService handlers.AuthService) *LoginPostHandler {
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

	session, err := h.authService.Login(r.Context(), email, password)
	if err != nil {
		slog.Error(err.Error())
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	utils.SetCookie(w, session.Token, "/home", session.Lifetime())
	w.Header().Set("HX-REDIRECT", "/home")
	w.WriteHeader(http.StatusOK)
}
