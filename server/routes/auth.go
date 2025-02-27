package routes

import (
	"github.com/go-chi/chi"
	"github.com/josuetorr/frequent-flyer/server/handlers/api"
)

func NewAuthRoutes(authService api.AuthService, userService api.UserService) chi.Router {
	r := chi.NewRouter()

	r.Post("/signup", api.Signup(authService).ServeHTTP())
	r.Post("/login", api.Login(authService).ServeHTTP())
	r.Post("/refresh", api.RefreshAccessToken(userService).ServeHTTP())

	return r
}
