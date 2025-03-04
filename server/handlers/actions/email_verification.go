package actions

import (
	"log/slog"
	"net/http"

	"github.com/go-chi/chi"
	au "github.com/josuetorr/frequent-flyer/internal/utils"
	"github.com/josuetorr/frequent-flyer/server/handlers"
	et "github.com/josuetorr/frequent-flyer/server/internal/utils/email_token"
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

	userID, err := et.VerifyToken(token, au.GetEmailVerificationSecret())
	switch err.(type) {
	}
	if err != nil {
		slog.Error(err.Error())
		http.Error(w, "Resource not found", http.StatusBadRequest)
		return
	}

	if err := h.userService.VerifyUser(r.Context(), userID); err != nil {
		slog.Error(err.Error())
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("HX-REDIRECT", "/home")
	w.WriteHeader(http.StatusNoContent)
}
