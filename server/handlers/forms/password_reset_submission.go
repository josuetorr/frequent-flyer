package forms

import (
	"errors"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/josuetorr/frequent-flyer/internal/utils"
	"github.com/josuetorr/frequent-flyer/server/handlers"
	"github.com/josuetorr/frequent-flyer/server/internal/utils/responder"
	"github.com/josuetorr/frequent-flyer/web/templates/components"
)

func HandlePasswordResetSubmission(userService handlers.UserService) responder.AppHandler {
	return func(w http.ResponseWriter, r *http.Request) *responder.AppError {
		token := chi.URLParam(r, "token")
		userId, err := utils.VerifyToken(token, utils.GetEmailSecret())
		if err != nil {
			return responder.NewNotFound(err, nil)
		}

		if err := r.ParseForm(); err != nil {
			return responder.NewBadRequest(err, components.AlertError("Invalid form"))
		}

		password := r.FormValue("password")
		confirmation := r.FormValue("password-confirm")

		const minLength = 8
		if len(password) < minLength {
			err := errors.New("Password must be at least 8 characters long")
			return responder.NewBadRequest(err, components.AlertError(err.Error()))
		}

		if password != confirmation {
			err := errors.New("Passwords do not match")
			return responder.NewBadRequest(err, components.AlertError(err.Error()))
		}

		if err := userService.UpdatePassword(r.Context(), userId, password); err != nil {
			return responder.NewInternalServer(err, components.AlertError("Oops... something went wrong"))
		}

		responder.NewAccepted(components.AlertSuccess("Password updated")).Respond(w, r)
		return nil
	}
}
