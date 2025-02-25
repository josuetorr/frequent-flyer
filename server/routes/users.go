package routes

import (
	"log/slog"

	"github.com/go-chi/chi"
	"github.com/josuetorr/frequent-flyer/server/data"
	"github.com/josuetorr/frequent-flyer/server/handlers"
	"github.com/josuetorr/frequent-flyer/server/services"
)

func NewUserRoutes(log *slog.Logger, db *data.DBPool) chi.Router {
	r := chi.NewRouter()

	userRepo := data.NewUserRepositor(db)
	userService := services.NewUserService(userRepo)

	r.Post("/", handlers.NewPostUserHandler(log, userService).ServeHTTP)
	r.Get("/{id}", handlers.NewGetUserHandler(log, userService).ServeHttp)
	r.Put("/{id}", handlers.NewPutUserHanlder(log, userService).ServeHttp)
	r.Delete("/{id}", handlers.NewDeleteUserHanlder().ServeHttp)

	return r
}
