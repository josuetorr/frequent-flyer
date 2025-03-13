package routes

import (
	"net/http"

	"github.com/go-chi/chi"
	cm "github.com/go-chi/chi/middleware"
	"github.com/josuetorr/frequent-flyer/internal/data"
	"github.com/josuetorr/frequent-flyer/internal/services"
	"github.com/josuetorr/frequent-flyer/internal/utils"
	"github.com/josuetorr/frequent-flyer/server/handlers"
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

	tokenSecret := utils.GetTokenSecret()
	shk := utils.GetSessionHashKey()
	sbk := utils.GetSessionBlockKey()

	r.Group(func(r chi.Router) {
		r.Use(authMw.RedirectIfLoggedIn)

		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			http.Redirect(w, r, handlers.LoginEndpoint, http.StatusFound)
		})
		r.Method("GET", handlers.LoginEndpoint, responder.AppHandler(pages.HandleLogin))
		r.Method("POST", handlers.LoginEndpoint, forms.HandleLoginForm(sessionCookieName, authService, shk, sbk))

		r.Method("GET", handlers.SignupEndpoint, responder.AppHandler(pages.HandleSignup))
		r.Method("POST", handlers.SignupEndpoint, forms.HandleSignupForm(authService, mailService))
	})

	r.Group(func(r chi.Router) {
		r.Use(authMw.Authorized)

		r.Method("GET", handlers.HomeEndpoint, responder.AppHandler(pages.HandleHome))
		r.Method("POST", handlers.LogoutEndpoint, forms.HandleLogout(sessionCookieName))
	})

	r.Method("GET", handlers.VerifyEmailEndpoint+"/{token}", actions.HandleEmailVerification(userService))
	r.Method("GET", handlers.PasswordResetEmailSubmissionEndpoint, responder.AppHandler(pages.HandlePasswordResetEmailSubmission))
	r.Method("POST", handlers.PasswordResetEmailSubmissionEndpoint, forms.HandlePasswordResetEmailSubmission(userService, mailService))
	r.Method("GET", handlers.PasswordResetEndpoint+"/{token}", pages.HandlePasswordResetSubmission(userService))
	r.Method("POST", handlers.PasswordResetEndpoint+"/{token}", forms.HandlePasswordResetSubmission(userService, tokenSecret))

	return r
}
