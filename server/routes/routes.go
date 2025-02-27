package routes

import (
	"github.com/go-chi/chi"
	cm "github.com/go-chi/chi/middleware"
	"github.com/josuetorr/frequent-flyer/server/data"
	m "github.com/josuetorr/frequent-flyer/server/middleware"
	"github.com/josuetorr/frequent-flyer/server/services"
	"github.com/josuetorr/frequent-flyer/server/utils"
)

func RegisterRoutes(db *data.DBPool) chi.Router {
	r := chi.NewRouter()
	r.Use(cm.Logger)

	userRepo := data.NewUserRepositor(db)
	userService := services.NewUserService(userRepo)
	authService := services.NewAuthService(userRepo)

	r.Route("/api/v1", func(r chi.Router) {
		r.Use(cm.AllowContentType(utils.ContentTypeJSON))

		r.Mount("/auth", NewAuthRoutes(authService, userRepo))
		r.Group(func(r chi.Router) {
			r.Use(m.AuthMiddleware)

			r.Mount("/users", NewUserRoutes(userService))
			r.Mount("/products", NewProductsRoutes())
			r.Mount("/stores", NewStoreRoutes())
		})
	})

	return r
}
