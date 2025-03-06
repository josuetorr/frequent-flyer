package forms

import (
	"net/http"

	"github.com/josuetorr/frequent-flyer/server/internal/utils/responder"
)

func HandlePasswordResetSubmission() responder.AppHandler {
	return func(w http.ResponseWriter, r *http.Request) *responder.AppError {
		return nil
	}
}
