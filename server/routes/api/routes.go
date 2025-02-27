package api

import (
	"github.com/go-chi/chi"
	cm "github.com/go-chi/chi/middleware"
	"github.com/josuetorr/frequent-flyer/server/handlers/api"
	m "github.com/josuetorr/frequent-flyer/server/middleware"
	"github.com/josuetorr/frequent-flyer/server/utils"
)

func RegisterApiRoutes(authService api.AuthService, userService api.UserService) chi.Router {
	r := chi.NewRouter()

	r.Use(cm.AllowContentType(utils.ContentTypeJSON))

	r.Mount("/auth", NewAuthRoutes(authService, userService))
	r.Group(func(r chi.Router) {
		r.Use(m.AuthMiddleware)

		r.Mount("/users", NewUserRoutes(userService))
		r.Mount("/products", NewProductsRoutes())
		r.Mount("/stores", NewStoreRoutes())
	})
	return r
}
