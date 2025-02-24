package handlers

import "net/http"

type PutUserHandler struct{}

func NewPutUserHanlder() PutUserHandler {
	return PutUserHandler{}
}

func (h PutUserHandler) ServeHttp(w http.ResponseWriter, r *http.Request) {}
