package actions

import (
	"errors"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi"
	au "github.com/josuetorr/frequent-flyer/internal/utils"
	emailToken "github.com/josuetorr/frequent-flyer/internal/utils/email_token"
	"github.com/josuetorr/frequent-flyer/server/handlers"
	"github.com/josuetorr/frequent-flyer/server/internal/utils/responder"
	emailTemplates "github.com/josuetorr/frequent-flyer/web/templates/email"
)

func HandleEmailVerification(userService handlers.UserService) responder.AppHandler {
	return func(w http.ResponseWriter, r *http.Request) *responder.AppError {
		token := chi.URLParam(r, "token")

		userID, err := emailToken.VerifyToken(token, au.GetEmailVerificationSecret())
		if err != nil {
			slog.Error(err.Error())
			switch {
			case errors.Is(err, emailToken.InvalidTokenErr):
			case errors.Is(err, emailToken.InvalidSignatureErr):
				return responder.NewBadRequest(err, nil, emailTemplates.Error("Invalid token"))
			case errors.Is(err, emailToken.ExpiredTokenErr):
				return responder.NewBadRequest(err, nil, emailTemplates.Error("Token expired"))
			default:
				return responder.NewInternalServer(err, nil, emailTemplates.Error("Internal server error"))
			}
		}

		if err := userService.VerifyUser(r.Context(), userID); err != nil {
			return responder.NewInternalServer(err, nil, nil)
		}

		http.Redirect(w, r, "/login", http.StatusSeeOther)
		w.WriteHeader(http.StatusNoContent)
		return nil
	}
}
