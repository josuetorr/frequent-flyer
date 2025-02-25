package handlers

import (
	"log/slog"
	"net/http"
	"strings"

	"github.com/go-chi/chi"
)

type DeleteUserHandler struct {
	log         *slog.Logger
	userService UserService
}

func NewDeleteUserHanlder(log *slog.Logger, service UserService) *DeleteUserHandler {
	return &DeleteUserHandler{log: log, userService: service}
}

func (h *DeleteUserHandler) ServeHttp(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	hard := strings.ToLower(r.URL.Query().Get("hard"))

	if err := h.userService.Delete(r.Context(), id, hard == "true"); err != nil {
		h.log.Error(err.Error())
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
