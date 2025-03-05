package actions

import (
	"errors"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi"
	au "github.com/josuetorr/frequent-flyer/internal/utils"
	emailToken "github.com/josuetorr/frequent-flyer/internal/utils/email_token"
	"github.com/josuetorr/frequent-flyer/server/handlers"
	emailTemplates "github.com/josuetorr/frequent-flyer/web/templates/email"
)

type EmailVerificationHandler struct {
	userService handlers.UserService
}

func NewEmailVerificationHandler(userService handlers.UserService) *EmailVerificationHandler {
	return &EmailVerificationHandler{
		userService: userService,
	}
}

func (h *EmailVerificationHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	token := chi.URLParam(r, "token")

	userID, err := emailToken.VerifyToken(token, au.GetEmailVerificationSecret())
	ctx := r.Context()
	if err != nil {
		slog.Error(err.Error())
		switch {
		case errors.Is(err, emailToken.InvalidTokenErr):
			emailTemplates.Error("Invalid token").Render(ctx, w)
		case errors.Is(err, emailToken.InvalidSignatureErr):
			emailTemplates.Error("Invalid token").Render(ctx, w)
		case errors.Is(err, emailToken.ExpiredTokenErr):
			emailTemplates.Error("Token expired").Render(ctx, w)
		default:
			emailTemplates.Error("Internal server error").Render(ctx, w)
		}
		return
	}

	if err := h.userService.VerifyUser(r.Context(), userID); err != nil {
		slog.Error(err.Error())
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/login", http.StatusSeeOther)
	w.WriteHeader(http.StatusNoContent)
}
