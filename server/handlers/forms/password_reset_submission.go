package forms

import (
	"errors"
	"net/http"

	"github.com/josuetorr/frequent-flyer/server/handlers"
	"github.com/josuetorr/frequent-flyer/server/internal/utils/responder"
	"github.com/josuetorr/frequent-flyer/web/templates/components"
)

func HandlePasswordResetSubmission(userService handlers.UserService) responder.AppHandler {
	return func(w http.ResponseWriter, r *http.Request) *responder.AppError {
		if err := r.ParseForm(); err != nil {
			return responder.NewBadRequest(err, components.AlertError("Invalid form"))
		}

		password := r.FormValue("password")
		confirmation := r.FormValue("password-confirmation")

		if password != confirmation {
			err := errors.New("Passwords do not match")
			return responder.NewBadRequest(err, components.AlertError(err.Error()))
		}

		return nil
	}
}
