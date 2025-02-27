package routes

import (
	"github.com/go-chi/chi"
	"github.com/josuetorr/frequent-flyer/server/handlers"
)

func NewAuthRoutes(authService handlers.AuthService, userService handlers.UserService) chi.Router {
	r := chi.NewRouter()

	r.Post("/signup", handlers.Signup(authService).ServeHTTP())
	r.Post("/login", handlers.Login(authService).ServeHTTP())
	r.Post("/refresh", handlers.RefreshAccessToken(userService).ServeHTTP())

	return r
}
