package pages

import (
	"net/http"

	"github.com/josuetorr/frequent-flyer/server/internal/utils/responder"
	"github.com/josuetorr/frequent-flyer/web/templates/pages"
)

func HandleLogin(w http.ResponseWriter, r *http.Request) *responder.AppError {
	responder.NewOk(pages.Login()).Respond(w, r)
	return nil
}
