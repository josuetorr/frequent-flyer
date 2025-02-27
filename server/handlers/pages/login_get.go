package pages

import (
	"log/slog"
	"net/http"

	"github.com/josuetorr/frequent-flyer/server/templates/pages"
)

type LoginPageHandler struct{}

func NewLoginPageHandler() *LoginPageHandler {
	return &LoginPageHandler{}
}

func (h *LoginPageHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := pages.Login().Render(r.Context(), w); err != nil {
		slog.Error(err.Error())
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}
