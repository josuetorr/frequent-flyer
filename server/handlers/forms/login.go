package forms

import (
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"net/mail"

	"github.com/josuetorr/frequent-flyer/internal/services"
	"github.com/josuetorr/frequent-flyer/internal/utils"
	"github.com/josuetorr/frequent-flyer/server/handlers"
	"github.com/josuetorr/frequent-flyer/server/internal/utils/responder"
	errorTempl "github.com/josuetorr/frequent-flyer/web/templates/errors"
)

func HandleLoginForm(sessionCookieName string, authService handlers.AuthService) responder.AppHandler {
	return func(w http.ResponseWriter, r *http.Request) *responder.AppError {
		if err := r.ParseForm(); err != nil {
			slog.Error(err.Error())
			return responder.NewBadRequest(err, nil)
		}

		email := r.FormValue("email")
		password := r.FormValue("password")

		if _, err := mail.ParseAddress(email); err != nil {
			w.Header().Set("HX-FOCUS", "#email")
			return responder.NewBadRequest(
				err,
				errorTempl.Alert("Invalid email"),
			)
		}

		session, err := authService.Login(r.Context(), email, password)
		if err != nil {
			switch {
			case errors.Is(err, services.InvalidCredentialError):
				return responder.NewBadRequest(err, errorTempl.Alert(err.Error()))
			default:
				slog.Error(err.Error())
				return responder.NewInternalServer(err, errorTempl.Alert("Oops... something whent wrong"))
			}
		}

		cookieValue := fmt.Sprintf("%s:%s", session.ID, session.UserID)
		encoded, err := utils.EncodeCookie(sessionCookieName, cookieValue)
		if err != nil {
			return responder.NewInternalServer(err, errorTempl.Alert("Oops... something whent wrong"))
		}

		// TODO: added HTTPs and Secure
		http.SetCookie(w, &http.Cookie{
			Name:  sessionCookieName,
			Value: encoded,
			// HttpOnly: true,
			// Secure:   true,
			Path:     "/",
			MaxAge:   session.Lifetime(),
			SameSite: http.SameSiteStrictMode,
		})

		w.Header().Set("HX-REDIRECT", "/home")
		w.WriteHeader(http.StatusOK)
		return nil
	}
}
