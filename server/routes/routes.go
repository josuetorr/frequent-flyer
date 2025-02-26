package routes

import (
	"log/slog"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/josuetorr/frequent-flyer/server/utils"
)

func RegisterRoutes(log *slog.Logger, db *DBPool) chi.Router {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Route("/api/v1", func(r chi.Router) {
		r.Use(middleware.AllowContentType(utils.ContentTypeJSON))

		r.Mount("/auth", NewAuthRoutes())
		r.Mount("/users", NewUserRoutes(log, db))
		r.Mount("/products", NewProductsRoutes())
		r.Mount("/stores", NewStoreRoutes())
	})

	return r
}
