package web

import (
	"net/http"

	"github.com/josuetorr/frequent-flyer/web/templates/pages"
)

type SignupPageHandler struct{}

func NewSignupPageHandler() *SignupPageHandler {
	return &SignupPageHandler{}
}

func (h *SignupPageHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	templates.Signup().Render(r.Context(), w)
}
