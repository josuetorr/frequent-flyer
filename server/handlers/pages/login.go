package pages

import (
	"net/http"

	"github.com/josuetorr/frequent-flyer/server/internal/utils/responder"
	"github.com/josuetorr/frequent-flyer/web/templates/pages"
)

type LoginPageHandler struct{}

func NewLoginPageHandler() *LoginPageHandler {
	return &LoginPageHandler{}
}

func HandleLoginPage(w http.ResponseWriter, r *http.Request) *responder.AppError {
	responder.NewOk(nil, templates.Login()).Respond(w, r)
	return nil
}
