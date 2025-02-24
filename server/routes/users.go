package routes

import (
	"github.com/go-chi/chi"
	"github.com/josuetorr/frequent-flyer/server/handlers"
)

func NewUserRoutes() chi.Router {
	r := chi.NewRouter()

	r.Post("/", handlers.NewPostUserHandler().ServeHTTP)

	return r
}
