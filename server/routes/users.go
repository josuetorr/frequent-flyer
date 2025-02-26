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
	r.Get("/{id}", handlers.GetUser(userService).ServeHTTP())
	r.Put("/{id}", handlers.NewPutUserHanlder(log, userService).ServeHTTP)
	r.Delete("/{id}", handlers.NewDeleteUserHanlder(log, userService).ServeHTTP)

	return r
}
