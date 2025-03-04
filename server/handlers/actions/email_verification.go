package actions

import (
	"log/slog"
	"net/http"

	"github.com/go-chi/chi"
	au "github.com/josuetorr/frequent-flyer/internal/utils"
	"github.com/josuetorr/frequent-flyer/server/handlers"
	su "github.com/josuetorr/frequent-flyer/server/internal/utils"
)

type EmailVerificationHandler struct {
	authService handlers.AuthService
}

func NewEmailVerificationHandler(authService handlers.AuthService) *EmailVerificationHandler {
	return &EmailVerificationHandler{
		authService: authService,
	}
}

func (h *EmailVerificationHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	token := chi.URLParam(r, "token")

	userID, err := su.VerifyToken(token, au.GetEmailVerificationSecret())
	if err != nil {
		slog.Error(err.Error())
		http.Error(w, "Resource not found", http.StatusBadRequest)
		return
	}

	if err := h.authService.VerifyUser(r.Context(), userID); err != nil {
		slog.Error(err.Error())
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("HX-REDIRECT", "/home")
	w.WriteHeader(http.StatusNoContent)
}
