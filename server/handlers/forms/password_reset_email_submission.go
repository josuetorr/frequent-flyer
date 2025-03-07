package forms

import (
	"net/http"
	"net/mail"

	"github.com/josuetorr/frequent-flyer/internal/utils"
	"github.com/josuetorr/frequent-flyer/server/handlers"
	"github.com/josuetorr/frequent-flyer/server/internal/utils/responder"
	errorTempl "github.com/josuetorr/frequent-flyer/web/templates/errors"
)

func HandlePasswordResetEmailSubmission(userService handlers.UserService, mailService handlers.MailService) responder.AppHandler {
	return func(w http.ResponseWriter, r *http.Request) *responder.AppError {
		if err := r.ParseForm(); err != nil {
			return responder.NewBadRequest(err, nil)
		}

		email := r.FormValue("email")

		if _, err := mail.ParseAddress(email); err != nil {
			return responder.NewBadRequest(err, errorTempl.Alert("Invalid email"))
		}

		ctx := r.Context()
		u, err := userService.GetByEmail(ctx, email)
		if err != nil || u == nil {
			return responder.NewBadRequest(err, errorTempl.Alert("User not found"))
		}

		link := mailService.GenerateEmailLink(u.ID, "password-reset", utils.GetEmailSecret())
		if err := mailService.SendPasswordResetEmail(ctx, link, email); err != nil {
			return responder.NewInternalServer(err, errorTempl.Alert("Oops... something went from whilst sending email"))
		}

		// TODO: move "Alert" to components instead of "errors"
		responder.NewOk(errorTempl.Alert("Email has been sent")).Respond(w, r)
		return nil
	}
}
