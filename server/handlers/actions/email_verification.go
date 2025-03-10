package actions

import (
	"errors"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/josuetorr/frequent-flyer/internal/utils"
	"github.com/josuetorr/frequent-flyer/server/handlers"
	"github.com/josuetorr/frequent-flyer/server/internal/utils/responder"
	"github.com/josuetorr/frequent-flyer/web/templates/components"
	emailTemplates "github.com/josuetorr/frequent-flyer/web/templates/email"
)

func HandleEmailVerification(userService handlers.UserService) responder.AppHandler {
	return func(w http.ResponseWriter, r *http.Request) *responder.AppError {
		token := chi.URLParam(r, "token")

		userID, err := utils.VerifyToken(token, utils.GetTokenSecret())
		if err != nil {
			switch {
			case errors.Is(err, utils.InvalidTokenErr),
				errors.Is(err, utils.InvalidSignatureErr):
				return responder.NewBadRequest(err, emailTemplates.Error("Invalid token"))
			case errors.Is(err, utils.ExpiredTokenErr):
				return responder.NewBadRequest(err, emailTemplates.Error("Token expired"))
			default:
				return responder.NewInternalServer(err, emailTemplates.Error("Internal server error"))
			}
		}

		u, err := userService.GetById(r.Context(), userID)
		if err != nil || u == nil {
			return responder.NewNotFound(err, components.AlertError("User not found"))
		}
		if err := userService.VerifyUser(r.Context(), userID); err != nil {
			return responder.NewInternalServer(err, nil)
		}

		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return nil
	}
}
