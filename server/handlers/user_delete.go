package handlers

import "net/http"

type DeleteUserHandler struct{}

func NewDeleteUserHanlder() *DeleteUserHandler {
	return &DeleteUserHandler{}
}

func (h *DeleteUserHandler) ServeHttp(w http.ResponseWriter, r *http.Request) {}
