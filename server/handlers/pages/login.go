package pages

import (
	"net/http"

	"github.com/josuetorr/frequent-flyer/web/templates/pages"
)

type LoginPageHandler struct{}

func NewLoginPageHandler() *LoginPageHandler {
	return &LoginPageHandler{}
}

func (h *LoginPageHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	templates.Login().Render(r.Context(), w)
}
