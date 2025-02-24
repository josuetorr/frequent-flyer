package routes

import (
	"github.com/go-chi/chi"
	"github.com/josuetorr/frequent-flyer/server/handlers"
)

func NewUserRoutes() chi.Router {
	r := chi.NewRouter()

	r.Post("/", handlers.NewPostUserHandler().ServeHTTP)
	r.Get("/{id}", handlers.NewGetUserHanlder().ServeHttp)
	r.Put("/{id}", handlers.NewPutUserHanlder().ServeHttp)
	r.Delete("/{id}", handlers.NewDeleteUserHanlder().ServeHttp)

	return r
}
