package routes

import (
	"github.com/go-chi/chi"
	cm "github.com/go-chi/chi/middleware"
	"github.com/josuetorr/frequent-flyer/internal/data"
	"github.com/josuetorr/frequent-flyer/internal/services"
	"github.com/josuetorr/frequent-flyer/server/routes/api"
	"github.com/josuetorr/frequent-flyer/server/routes/pages"
)

func RegisterRoutes(db *data.DBPool) chi.Router {
	r := chi.NewRouter()
	r.Use(cm.Logger)

	userRepo := data.NewUserRepositor(db)
	userService := services.NewUserService(userRepo)
	authService := services.NewAuthService(userRepo)

	r.Mount("/api/v1", api.RegisterApiRoutes(authService, userService))
	r.Mount("/", pages.RegisterPagesRoutes())

	return r
}
