package handlers

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi"
)

type GetUserHandler struct {
	log         *slog.Logger
	userService UserService
}

func NewGetUserHandler(log *slog.Logger, service UserService) *GetUserHandler {
	return &GetUserHandler{log: log, userService: service}
}

func (h *GetUserHandler) ServeHttp(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	u, err := h.userService.Get(r.Context(), id)
	if err != nil {
		h.log.Error(err.Error())
		http.Error(w, "Resource not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&u)
}
