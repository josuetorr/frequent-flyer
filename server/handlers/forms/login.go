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
	"github.com/josuetorr/frequent-flyer/web/templates/components"
)

func HandleLoginForm(
	sessionCookieName string,
	authService handlers.AuthService,
	sessionHashKey string,
	sessionBlockKey string,
) responder.AppHandler {
	return func(w http.ResponseWriter, r *http.Request) *responder.AppError {
		if r.Header.Get("Content-Type") != "application/x-www-form-urlencoded" {
			return responder.NewUnsupportedMediaType(errors.New("Unsupported media type"), nil)
		}
		if err := r.ParseForm(); err != nil {
			slog.Error(err.Error())
			return responder.NewBadRequest(err, nil)
		}

		email := r.FormValue("email")
		password := r.FormValue("password")

		if _, err := mail.ParseAddress(email); err != nil {
			w.Header().Set("HX-FOCUS", "#email")
			return responder.NewBadRequest(err, components.AlertError("Invalid email"))
		}

		session, err := authService.Login(r.Context(), email, password)
		if err != nil {
			switch {
			case errors.Is(err, services.InvalidCredentialError):
				return responder.NewBadRequest(err, components.AlertError(err.Error()))
			default:
				slog.Error(err.Error())
				return responder.NewInternalServer(err, components.AlertError("Oops... something whent wrong"))
			}
		}

		cookieValue := fmt.Sprintf("%s:%s", session.ID, session.UserID)
		encoded, err := utils.EncodeCookie(sessionCookieName, cookieValue, sessionHashKey, sessionBlockKey)
		if err != nil {
			return responder.NewInternalServer(err, components.AlertError("Oops... something whent wrong"))
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

		w.Header().Set("HX-REDIRECT", handlers.HomeEndpoint)
		w.WriteHeader(http.StatusOK)
		return nil
	}
}
