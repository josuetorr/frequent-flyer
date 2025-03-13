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
		cookie := &http.Cookie{
			Name:  sessionCookieName,
			Value: "",
			// HttpOnly: true,
			// Secure:   true,
			Path:     "/",
			MaxAge:   -1,
			SameSite: http.SameSiteStrictMode,
		}
		http.SetCookie(w, cookie)
		w.Header().Set("HX-REDIRECT", handlers.LoginEndpoint)
		w.WriteHeader(http.StatusOK)

		return nil
	}
}
