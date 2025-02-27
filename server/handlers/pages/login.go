package pages

import (
	"net/http"
	// "github.com/josuetorr/web/templates"
)

type LoginPageHandler struct{}

func NewLoginPageHandler() *LoginPageHandler {
	return &LoginPageHandler{}
}

func (h *LoginPageHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// t := templates
}
