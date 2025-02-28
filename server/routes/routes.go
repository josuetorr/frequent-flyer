package routes

import (
	"github.com/go-chi/chi"
	cm "github.com/go-chi/chi/middleware"
	"github.com/josuetorr/frequent-flyer/internal/data"
	"github.com/josuetorr/frequent-flyer/server/routes/web"
)

func RegisterRoutes(db *data.DBPool) chi.Router {
	r := chi.NewRouter()
	r.Use(cm.Logger)

	r.Mount("/", pages.RegisterWebRoutes())

	return r
}
