package forms

import (
	"log/slog"
	"net/http"

	"github.com/josuetorr/frequent-flyer/server/handlers"
	"github.com/josuetorr/frequent-flyer/server/internal/utils"
)

type LogoutPostHandler struct {
	authService handlers.AuthService
}

func NewLogoutHandler(authService handlers.AuthService) *LogoutPostHandler {
	return &LogoutPostHandler{authService: authService}
}

func (h *LogoutPostHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie(utils.SessionCookieName)
	if err != nil {
		slog.Error(err.Error())
		http.Error(w, "No session found", http.StatusUnauthorized)
		return
	}

	if err := h.authService.Logout(r.Context(), cookie.Value); err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	utils.InvalidateCookie(w)
	w.Header().Set("HX-REDIRECT", "/login")
}
