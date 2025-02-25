package routes

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func RegisterRoutes(db *DBPool) chi.Router {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Route("/api/v1", func(r chi.Router) {
		r.Mount("/users", NewUserRoutes(db))
		r.Mount("/products", NewProductsRoutes())
		r.Mount("/stores", NewStoreRoutes())
	})

	return r
}
