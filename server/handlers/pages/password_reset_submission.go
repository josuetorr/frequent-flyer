package pages

import (
	"errors"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/josuetorr/frequent-flyer/internal/utils"
	"github.com/josuetorr/frequent-flyer/internal/utils/email_token"
	"github.com/josuetorr/frequent-flyer/server/handlers"
	"github.com/josuetorr/frequent-flyer/server/internal/utils/responder"
	"github.com/josuetorr/frequent-flyer/web/templates/components"
	emailTemplates "github.com/josuetorr/frequent-flyer/web/templates/email"
	"github.com/josuetorr/frequent-flyer/web/templates/pages"
)

func HandlePasswordResetSubmission(userService handlers.UserService) responder.AppHandler {
	return func(w http.ResponseWriter, r *http.Request) *responder.AppError {
		token := chi.URLParam(r, "token")

		userID, err := emailtoken.VerifyToken(token, utils.GetEmailSecret())
		if err != nil {
			switch {
			case errors.Is(err, emailtoken.InvalidTokenErr),
				errors.Is(err, emailtoken.InvalidSignatureErr):
				return responder.NewBadRequest(err, emailTemplates.Error("Invalid token"))
			case errors.Is(err, emailtoken.ExpiredTokenErr):
				return responder.NewBadRequest(err, emailTemplates.Error("Token expired"))
			default:
				return responder.NewInternalServer(err, emailTemplates.Error("Internal server error"))
			}
		}

		u, err := userService.GetById(r.Context(), userID)
		if err != nil || u == nil {
			return responder.NewNotFound(err, components.AlertError("User not found"))
		}

		responder.NewOk(pages.PasswordReset(token)).Respond(w, r)
		return nil
	}
}
