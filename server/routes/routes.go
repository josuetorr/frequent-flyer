package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func RegisterRoutes() chi.Router {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Mount("/users", NewUserRoutes())
	r.Mount("/products", NewProductsRoutes())
	r.Mount("/stores", NewStoreRoutes())

	return r
}
