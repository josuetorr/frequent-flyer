package routes

import (
	"github.com/go-chi/chi"
	"github.com/josuetorr/frequent-flyer/server/handlers"
	"github.com/josuetorr/frequent-flyer/server/services"
)

func NewAuthRoutes(userRepo services.UserRepository) chi.Router {
	r := chi.NewRouter()

	authService := services.NewAuthService(userRepo)

	r.Post("/signup", handlers.Signup(authService).ServeHTTP())
	r.Post("/login", handlers.Login(authService).ServeHTTP())
	r.Post("/refresh", handlers.RefreshAccessToken().ServeHTTP())

	return r
}
