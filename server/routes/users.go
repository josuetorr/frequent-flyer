package routes

import (
	"github.com/go-chi/chi"
	"github.com/josuetorr/frequent-flyer/server/data"
	"github.com/josuetorr/frequent-flyer/server/handlers"
	"github.com/josuetorr/frequent-flyer/server/services"
)

func NewUserRoutes(db *data.DBPool) chi.Router {
	r := chi.NewRouter()

	userRepo := data.NewUserRepositor(db)
	userService := services.NewUserService(userRepo)

	r.Post("/", handlers.NewPostUserHandler(userService).ServeHTTP)
	r.Get("/{id}", handlers.NewGetUserHanlder().ServeHttp)
	r.Put("/{id}", handlers.NewPutUserHanlder().ServeHttp)
	r.Delete("/{id}", handlers.NewDeleteUserHanlder().ServeHttp)

	return r
}
