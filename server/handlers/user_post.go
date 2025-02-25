package handlers

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
)

type PostUserHandler struct {
	log         *slog.Logger
	userService UserService
}

func NewPostUserHandler(log *slog.Logger, service UserService) PostUserHandler {
	return PostUserHandler{log: log, userService: service}
}

func (h PostUserHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var u User
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		http.Error(w, "Invalid payload", http.StatusBadRequest)
		return
	}

	fmt.Printf("user: %+v\n", u)
	err := h.userService.Insert(r.Context(), &u)
	if err != nil {
		h.log.Error(err.Error())
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
