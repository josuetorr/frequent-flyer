package routes

import (
	"log/slog"

	"github.com/go-chi/chi"
	cm "github.com/go-chi/chi/middleware"
	"github.com/josuetorr/frequent-flyer/server/data"
	m "github.com/josuetorr/frequent-flyer/server/middleware"
	"github.com/josuetorr/frequent-flyer/server/utils"
)

func RegisterRoutes(log *slog.Logger, db *data.DBPool) chi.Router {
	r := chi.NewRouter()
	r.Use(cm.Logger)

	userRepo := data.NewUserRepositor(db)

	r.Route("/api/v1", func(r chi.Router) {
		r.Use(cm.AllowContentType(utils.ContentTypeJSON))

		r.Mount("/auth", NewAuthRoutes(userRepo))
		r.Group(func(r chi.Router) {
			r.Use(m.AuthMiddleware)

			r.Mount("/users", NewUserRoutes(userRepo))
			r.Mount("/products", NewProductsRoutes())
			r.Mount("/stores", NewStoreRoutes())
		})
	})

	return r
}
