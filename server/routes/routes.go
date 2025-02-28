package routes

import (
	"net/http"

	"github.com/go-chi/chi"
	cm "github.com/go-chi/chi/middleware"
	"github.com/josuetorr/frequent-flyer/internal/data"
	"github.com/josuetorr/frequent-flyer/internal/services"
	"github.com/josuetorr/frequent-flyer/server/handlers/forms"
	"github.com/josuetorr/frequent-flyer/server/handlers/pages"
	"github.com/josuetorr/frequent-flyer/server/internal/middleware"
)

func RegisterRoutes(db *data.DBPool) chi.Router {
	r := chi.NewRouter()
	r.Use(cm.Logger)

	userRepo := data.NewUserRepository(db)
	sessionRepo := data.NewSessionRepository(db)

	authService := services.NewAuthService(userRepo, sessionRepo)
	sessionService := services.NewSessionService(sessionRepo)

	r.Group(func(r chi.Router) {
		r.Use(middleware.RedirectIfLogged(sessionService))

		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			http.Redirect(w, r, "/login", http.StatusFound)
		})
		r.Method("GET", "/login", pages.NewLoginPageHandler())
		r.Method("POST", "/login", forms.NewLoginHandler(authService))

		r.Method("GET", "/signup", pages.NewSignupPageHandler())
		r.Method("POST", "/signup", forms.NewSignupHandler(authService))
	})

	r.Group(func(r chi.Router) {
		r.Use(middleware.AuthMiddlerware(sessionService))

		r.Method("GET", "/home", pages.NewHomePageHandler())
		r.Method("POST", "/logout", forms.NewLogoutHandler(authService))
	})

	return r
}
