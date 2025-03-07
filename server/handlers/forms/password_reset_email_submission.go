package forms

import (
	"net/http"
	"net/mail"

	"github.com/josuetorr/frequent-flyer/internal/utils"
	"github.com/josuetorr/frequent-flyer/server/handlers"
	"github.com/josuetorr/frequent-flyer/server/internal/utils/responder"
	"github.com/josuetorr/frequent-flyer/web/templates/components"
)

func HandlePasswordResetEmailSubmission(userService handlers.UserService, mailService handlers.MailService) responder.AppHandler {
	return func(w http.ResponseWriter, r *http.Request) *responder.AppError {
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
			return responder.NewBadRequest(err, components.AlertError("User not found"))
		}

		link := mailService.GenerateEmailLink(u.ID, "password-reset", utils.GetEmailSecret())
		if err := mailService.SendPasswordResetEmail(ctx, link, email); err != nil {
			return responder.NewInternalServer(err, components.AlertError("Oops... something went wrong whilst sending email"))
		}

		// TODO: move "Alert" to components instead of "errors"
		responder.NewOk(components.AlertSuccess("Email has been sent")).Respond(w, r)
		return nil
	}
}
