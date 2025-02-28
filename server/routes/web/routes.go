package pages

import (
	"github.com/go-chi/chi"
	"github.com/josuetorr/frequent-flyer/server/handlers/web"
)

func RegisterWebRoutes() chi.Router {
	r := chi.NewRouter()

	r.Method("GET", "/login", web.NewLoginPageHandler())
	r.Method("POST", "/login", web.NewLoginPostHandler())

	r.Method("GET", "/signup", web.NewSignupPageHandler())
	r.Method("GET", "/home", web.NewHomePageHandler())

	return r
}
