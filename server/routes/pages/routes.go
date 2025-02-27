package pages

import (
	"github.com/go-chi/chi"
	"github.com/josuetorr/frequent-flyer/server/handlers/pages"
)

func RegisterPagesRoutes() chi.Router {
	r := chi.NewRouter()

	r.Get("/", pages.NewLoginPageHandler().ServeHTTP)

	return r
}
