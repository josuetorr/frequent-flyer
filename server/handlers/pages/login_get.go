package pages

import "net/http"

type LoginPageHandler struct{}

func NewLoginPageHandler() *LoginPageHandler {
	return &LoginPageHandler{}
}

func (h *LoginPageHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
}
