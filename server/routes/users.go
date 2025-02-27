package routes

import (
	"github.com/go-chi/chi"
	"github.com/josuetorr/frequent-flyer/server/handlers/api"
)

func NewUserRoutes(userService api.UserService) chi.Router {
	r := chi.NewRouter()

	r.Post("/", api.CreateUser(userService).ServeHTTP())
	r.Get("/{id}", api.GetUser(userService).ServeHTTP())
	r.Put("/{id}", api.UpdateUser(userService).ServeHTTP())
	r.Delete("/{id}", api.DeleteUser(userService).ServeHTTP())

	return r
}
