package handlers

import (
	"encoding/json"
	"net/http"
)

type PostUserHandler struct{}

func NewPostUserHandler() PostUserHandler {
	return PostUserHandler{}
}

func (h PostUserHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ct := w.Header().Get("Content-Type")
	println(ct)
	data := map[string]any{"hello": "hello"}
	json.NewEncoder(w).Encode(data)
	w.WriteHeader(http.StatusOK)
}
