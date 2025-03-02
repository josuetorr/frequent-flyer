package forms

import (
	"net/http"

	"github.com/josuetorr/frequent-flyer/server/handlers"
)

type LogoutPostHandler struct {
	sessionCookieName string
	authService       handlers.AuthService
}

func NewLogoutHandler(sessionCookieName string, authService handlers.AuthService) *LogoutPostHandler {
	return &LogoutPostHandler{sessionCookieName: sessionCookieName, authService: authService}
}

// NOTE: Let's keep the session. This way when the user logs back, we can fetch their
// session. If they decide to kill their session, then we delete it.
func (h *LogoutPostHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:  h.sessionCookieName,
		Value: "",
		// HttpOnly: true,
		// Secure:   true,
		Path:     "/",
		MaxAge:   -1,
		SameSite: http.SameSiteStrictMode,
	})
	w.Header().Set("HX-REDIRECT", "/login")
}
