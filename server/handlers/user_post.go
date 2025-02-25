package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type PostUserHandler struct {
	userService UserService
}

func NewPostUserHandler(service UserService) PostUserHandler {
	return PostUserHandler{userService: service}
}

func (h PostUserHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var u User
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		http.Error(w, "Invalid payload", http.StatusBadRequest)
		return
	}

	fmt.Printf("user: %+v\n", u)
	w.WriteHeader(http.StatusCreated)
}
