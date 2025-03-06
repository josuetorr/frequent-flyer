package routes

import (
	"net/http"

	"github.com/go-chi/chi"
	cm "github.com/go-chi/chi/middleware"
	"github.com/josuetorr/frequent-flyer/internal/data"
	"github.com/josuetorr/frequent-flyer/internal/services"
	"github.com/josuetorr/frequent-flyer/server/handlers/actions"
	"github.com/josuetorr/frequent-flyer/server/handlers/forms"
	"github.com/josuetorr/frequent-flyer/server/handlers/pages"
	"github.com/josuetorr/frequent-flyer/server/internal/middleware"
	"github.com/josuetorr/frequent-flyer/server/internal/utils/responder"
)

func RegisterRoutes(db *data.DBPool) chi.Router {
	r := chi.NewRouter()
	r.Use(cm.Logger)

	sessionCookieName := "session_cookie"

	userRepo := data.NewUserRepository(db)
	sessionRepo := data.NewSessionRepository(db)

	userService := services.NewUserService(userRepo)
	authService := services.NewAuthService(userRepo, sessionRepo)
	sessionService := services.NewSessionService(sessionRepo)
	mailService := services.NewMailService()

	authMw := middleware.NewAuthMiddleware(sessionCookieName, sessionService)

	r.Group(func(r chi.Router) {
		r.Use(authMw.RedirectIfLoggedIn)

		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			http.Redirect(w, r, "/login", http.StatusFound)
		})
		r.Method("GET", "/login", responder.AppHandler(pages.HandleLoginPage))
		r.Method("POST", "/login", forms.HandleLoginForm(sessionCookieName, authService))

		r.Method("GET", "/signup", responder.AppHandler(pages.HandleSignupPage))
		r.Method("POST", "/signup", forms.HandleSignupForm(authService, mailService))
	})

	r.Group(func(r chi.Router) {
		r.Use(authMw.Authorized)

		r.Method("GET", "/home", responder.AppHandler(pages.HandleHomePage))
		r.Method("POST", "/logout", forms.HandleLogout(sessionCookieName))
	})

	r.Method("GET", "/verify-email/{token}", actions.HandleEmailVerification(userService))

	return r
}
