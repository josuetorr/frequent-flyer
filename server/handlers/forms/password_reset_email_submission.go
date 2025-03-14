package forms

import (
	"errors"
	"net/http"
	"net/mail"

	"github.com/josuetorr/frequent-flyer/internal/utils"
	"github.com/josuetorr/frequent-flyer/server/handlers"
	"github.com/josuetorr/frequent-flyer/server/internal/utils/responder"
	"github.com/josuetorr/frequent-flyer/web/templates/components"
)

func HandlePasswordResetEmailSubmission(userService handlers.UserService, mailService handlers.MailService) responder.AppHandler {
	return func(w http.ResponseWriter, r *http.Request) *responder.AppError {
		if r.Header.Get("Content-Type") != "application/x-www-form-urlencoded" {
			err := errors.New("Unsupported Media type")
			return responder.NewUnsupportedMediaType(err, nil)
		}
		if err := r.ParseForm(); err != nil {
			return responder.NewBadRequest(err, nil)
		}

		email := r.FormValue("email")

		if _, err := mail.ParseAddress(email); err != nil {
			return responder.NewBadRequest(err, components.AlertError("Invalid email"))
		}

		ctx := r.Context()
		u, err := userService.GetByEmail(ctx, email)
		if err != nil || u == nil {
			return responder.NewNotFound(err, components.AlertError("User not found"))
		}

		link := mailService.GenerateEmailLink(u.ID, "password-reset", utils.GetTokenSecret())
		if err := mailService.SendPasswordResetEmail(ctx, link, email); err != nil {
			return responder.NewInternalServer(err, components.AlertError("Oops... something went wrong whilst sending email"))
		}

		responder.NewOk(components.AlertSuccess("Email has been sent")).Respond(w, r)
		return nil
	}
}
