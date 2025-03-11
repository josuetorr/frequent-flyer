package forms

import (
	"net/http"

	"github.com/josuetorr/frequent-flyer/server/handlers"
	"github.com/josuetorr/frequent-flyer/server/internal/utils/responder"
)

// NOTE: Let's keep the session. This way when the user logs back, we can fetch their
// session. If they decide to kill their session, then we delete it.
func HandleLogout(sessionCookieName string) responder.AppHandler {
	return func(w http.ResponseWriter, r *http.Request) *responder.AppError {
		http.SetCookie(w, &http.Cookie{
			Name:  sessionCookieName,
			Value: "",
			// HttpOnly: true,
			// Secure:   true,
			Path:     "/",
			MaxAge:   -1,
			SameSite: http.SameSiteStrictMode,
		})
		w.Header().Set("HX-REDIRECT", handlers.LoginEndpoint)

		return nil
	}
}
