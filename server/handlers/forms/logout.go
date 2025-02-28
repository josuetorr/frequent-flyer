package forms

import (
	"net/http"

	"github.com/josuetorr/frequent-flyer/server/handlers"
)

type LogoutPostHandler struct {
	authService handlers.AuthService
}

func NewLogoutHandler(authService handlers.AuthService) *LogoutPostHandler {
	return &LogoutPostHandler{authService: authService}
}

func (h *LogoutPostHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {}
