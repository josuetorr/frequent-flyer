package handlers

import "net/http"

type GetUserHandler struct{}

func NewGetUserHanlder() GetUserHandler {
	return GetUserHandler{}
}

func (h GetUserHandler) ServeHttp(w http.ResponseWriter, r *http.Request) {}
