package handlers

import "net/http"

type PostUserHandler struct{}

func NewPostUserHandler() *PostUserHandler {
	return &PostUserHandler{}
}

func (h *PostUserHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {}
