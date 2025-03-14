package forms

import (
	"errors"
	"fmt"
	"net/http"
	"net/mail"

	"github.com/josuetorr/frequent-flyer/internal/services"
	"github.com/josuetorr/frequent-flyer/server/handlers"
	"github.com/josuetorr/frequent-flyer/server/internal/utils/responder"
	"github.com/josuetorr/frequent-flyer/web/templates/components"
)

func HandleSignupForm(
	authService handlers.AuthService,
	mailService handlers.MailService,
	tokenSecret string,
) responder.AppHandler {
	return func(w http.ResponseWriter, r *http.Request) *responder.AppError {
		if r.Header.Get("Content-Type") != "application/x-www-form-urlencoded" {
			err := errors.New("Unsupported Media type")
			return responder.NewUnsupportedMediaType(err, nil)
		}
		if err := r.ParseForm(); err != nil {
			return responder.NewBadRequest(err, nil)
		}

		email := r.FormValue("email")
		password := r.FormValue("password")
		passwordConfirm := r.FormValue("password-confirm")

		if _, err := mail.ParseAddress(email); err != nil {
			w.Header().Set("HX-FOCUS", "#email")
			return responder.NewBadRequest(err, components.AlertError("Invalid email"))
		}

		const minPasswordLen = 8
		if len(password) < minPasswordLen {
			err := errors.New(fmt.Sprintf("Password must be at least %d characters long", minPasswordLen))
			w.Header().Set("HX-FOCUS", "#password")
			return responder.NewBadRequest(err, components.AlertError(err.Error()))
		}

		if password != passwordConfirm {
			err := errors.New("Passwords do not match")
			w.Header().Set("HX-FOCUS", "#password-confirm")
			return responder.NewBadRequest(err, components.AlertError(err.Error()))
		}

		ctx := r.Context()
		userID, err := authService.Signup(ctx, email, password)
		if err != nil {
			switch {
			case errors.Is(err, services.UserAlreadyExistsError):
				return responder.NewBadRequest(err, components.AlertError(err.Error()))
			default:
				return responder.NewInternalServer(err, components.AlertError("Oops... something went wrong"))
			}
		}

		link := mailService.GenerateEmailLink(userID, handlers.VerifyEmailEndpoint, tokenSecret)

		if err := mailService.SendVerificationEmail(ctx, link, email); err != nil {
			return responder.NewInternalServer(err, components.AlertError("Oops... something went wrong"))
		}

		w.Header().Set("HX-REDIRECT", handlers.LoginEndpoint)
		w.WriteHeader(http.StatusCreated)
		return nil
	}
}
