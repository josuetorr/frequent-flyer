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

	sessionCookieName := "session_cookie"

	userRepo := data.NewUserRepository(db)
	sessionRepo := data.NewSessionRepository(db)

	authService := services.NewAuthService(userRepo, sessionRepo)
	sessionService := services.NewSessionService(sessionRepo)
	mailService := services.NewMailService()

	authMw := middleware.NewAuthMiddleware(sessionCookieName, sessionService)

	r.Group(func(r chi.Router) {
		r.Use(authMw.RedirectIfLogged)

		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			http.Redirect(w, r, "/login", http.StatusFound)
		})
		r.Method("GET", "/login", pages.NewLoginPageHandler())
		r.Method("POST", "/login", forms.NewLoginHandler(sessionCookieName, authService))

		r.Method("GET", "/signup", pages.NewSignupPageHandler())
		r.Method("POST", "/signup", forms.NewSignupHandler(authService, mailService))
	})

	r.Group(func(r chi.Router) {
		r.Use(authMw.Authorized)

		r.Method("GET", "/home", pages.NewHomePageHandler())
		r.Method("POST", "/logout", forms.NewLogoutHandler(sessionCookieName, authService))
	})

	return r
}
