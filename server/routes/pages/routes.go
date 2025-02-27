package pages

import (
	"github.com/go-chi/chi"
	"github.com/josuetorr/frequent-flyer/server/handlers/pages"
)

func RegisterPagesRoutes() chi.Router {
	r := chi.NewRouter()

	r.Get("/login", pages.NewLoginPageHandler().ServeHTTP)
	r.Get("/signup", pages.NewSignupPageHandler().ServeHTTP)
	r.Get("/home", pages.NewHomePageHandler().ServeHTTP)

	return r
}
