package handlers

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi"
)

type PutUserHandler struct {
	log         *slog.Logger
	userService UserService
}

func NewPutUserHanlder(log *slog.Logger, service UserService) *PutUserHandler {
	return &PutUserHandler{log: log, userService: service}
}

func (h *PutUserHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ct := r.Header.Get("Content-Type")
	if ct != "application/json" {
		http.Error(w, "Unsupported Media Type", http.StatusUnsupportedMediaType)
		return
	}

	id := chi.URLParam(r, "id")
	defer r.Body.Close()

	var updatedUser User
	if err := json.NewDecoder(r.Body).Decode(&updatedUser); err != nil {
		h.log.Error(err.Error())
		http.Error(w, "Invalid json", http.StatusBadRequest)
		return
	}

	if err := h.userService.Update(r.Context(), id, &updatedUser); err != nil {
		h.log.Error(err.Error())
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
