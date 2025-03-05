package forms

import (
	"fmt"
	"log/slog"
	"net/http"
	"net/mail"

	"github.com/josuetorr/frequent-flyer/internal/utils"
	"github.com/josuetorr/frequent-flyer/server/handlers"
	"github.com/josuetorr/frequent-flyer/web/templates/errors"
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

	ctx := r.Context()
	email := r.FormValue("email")
	password := r.FormValue("password")

	_, err = mail.ParseAddress(email)
	if err != nil {
		w.Header().Set("HX-FOCUS", "email")
		w.WriteHeader(http.StatusBadRequest)
		errors.Alert("Invalid email").Render(ctx, w)
		return
	}

	session, err := h.authService.Login(r.Context(), email, password)
	if err != nil {
		slog.Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		errors.Alert("Oops... something went wrong").Render(ctx, w)
		return
	}

	cookieValue := fmt.Sprintf("%s:%s", session.ID, session.UserID)
	encoded, err := utils.EncodeCookie(h.sessionCookieName, cookieValue)
	if err != nil {
		slog.Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		errors.Alert("Oops... something went wrong").Render(ctx, w)
		return
	}

	// TODO: added HTTPs and Secure
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
