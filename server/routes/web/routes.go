package pages

import (
	"github.com/go-chi/chi"
	"github.com/josuetorr/frequent-flyer/server/handlers/web"
)

func RegisterWebRoutes() chi.Router {
	r := chi.NewRouter()

	r.Get("/login", web.NewLoginPageHandler().ServeHTTP)
	r.Get("/signup", web.NewSignupPageHandler().ServeHTTP)
	r.Get("/home", web.NewHomePageHandler().ServeHTTP)

	return r
}
