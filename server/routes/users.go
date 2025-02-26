package routes

import (
	"github.com/go-chi/chi"
	"github.com/josuetorr/frequent-flyer/server/handlers"
	"github.com/josuetorr/frequent-flyer/server/services"
)

func NewUserRoutes(userRepo services.UserRepository) chi.Router {
	r := chi.NewRouter()

	userService := services.NewUserService(userRepo)

	r.Post("/", handlers.CreateUser(userService).ServeHTTP())
	r.Get("/{id}", handlers.GetUser(userService).ServeHTTP())
	r.Put("/{id}", handlers.UpdateUser(userService).ServeHTTP())
	r.Delete("/{id}", handlers.DeleteUser(userService).ServeHTTP())

	return r
}
