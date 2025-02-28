package routes

import (
	"github.com/go-chi/chi"
	cm "github.com/go-chi/chi/middleware"
	"github.com/josuetorr/frequent-flyer/internal/data"
	"github.com/josuetorr/frequent-flyer/internal/services"
	"github.com/josuetorr/frequent-flyer/server/handlers/forms"
	"github.com/josuetorr/frequent-flyer/server/handlers/pages"
)

func RegisterRoutes(db *data.DBPool) chi.Router {
	r := chi.NewRouter()
	r.Use(cm.Logger)

	userRepo := data.NewUserRepository(db)
	sessionRepo := data.NewSessionRepository(db)

	authService := services.NewAuthService(userRepo, sessionRepo)

	r.Method("GET", "/login", pages.NewLoginPageHandler())
	r.Method("POST", "/login", forms.NewLoginHandler(authService))

	r.Method("GET", "/signup", pages.NewSignupPageHandler())
	r.Method("POST", "/signup", forms.NewSignupHandler(authService))

	r.Method("POST", "/logout", forms.NewLogoutHandler(authService))

	r.Method("GET", "/home", pages.NewHomePageHandler())

	return r
}
