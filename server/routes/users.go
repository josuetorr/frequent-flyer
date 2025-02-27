package routes

import (
	"github.com/go-chi/chi"
	"github.com/josuetorr/frequent-flyer/server/handlers/api"
)

func NewUserRoutes(userService handlers.UserService) chi.Router {
	r := chi.NewRouter()

	r.Post("/", handlers.CreateUser(userService).ServeHTTP())
	r.Get("/{id}", handlers.GetUser(userService).ServeHTTP())
	r.Put("/{id}", handlers.UpdateUser(userService).ServeHTTP())
	r.Delete("/{id}", handlers.DeleteUser(userService).ServeHTTP())

	return r
}
