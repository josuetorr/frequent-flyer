package forms

import (
	"fmt"
	"log/slog"
	"net/http"
	"net/mail"

	"github.com/josuetorr/frequent-flyer/internal/utils"
	"github.com/josuetorr/frequent-flyer/server/handlers"
)

type LoginPostHandler struct {
	sessionCookieName string
	authService       handlers.AuthService
}

func NewLoginHandler(sessionCookieName string, authService handlers.AuthService) *LoginPostHandler {
	return &LoginPostHandler{sessionCookieName: sessionCookieName, authService: authService}
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

	cookieValue := fmt.Sprintf("%s:%s", session.ID, session.UserID)
	encoded, err := utils.EncodeCookie(h.sessionCookieName, cookieValue)
	if err != nil {
		slog.Error(err.Error())
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:  h.sessionCookieName,
		Value: encoded,
		// HttpOnly: true,
		// Secure:   true,
		Path:     "/",
		MaxAge:   session.Lifetime(),
		SameSite: http.SameSiteStrictMode,
	})
	w.Header().Set("HX-REDIRECT", "/home")
	w.WriteHeader(http.StatusOK)
}
