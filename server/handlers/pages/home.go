package pages

import (
	"net/http"

	"github.com/josuetorr/frequent-flyer/web/templates/pages"
)

type HomePageHandler struct{}

func NewHomePageHandler() *HomePageHandler {
	return &HomePageHandler{}
}

func (h *HomePageHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	templates.Home().Render(r.Context(), w)
}
