package middleware

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"strings"

	"github.com/josuetorr/frequent-flyer/server/handlers"
)

type LoggedUser = string

type LoggedUserKey = LoggedUser("logger_user")

type Middleware = func(http.Handler) http.Handler

type AuthMiddleware struct {
	sessionCookieName string
	sessionService    handlers.SessionService
}

func NewAuthMiddleware(sessionCookieName string, sessionService handlers.SessionService) *AuthMiddleware {
	return &AuthMiddleware{sessionCookieName: sessionCookieName, sessionService: sessionService}
}

func (m *AuthMiddleware) Authorized(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		sessionCookie, err := r.Cookie(m.sessionCookieName)
		if err != nil {
			slog.Error(err.Error())
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}

    // TODO: decrypt session cookie
		values := strings.Split(sessionCookie.Value, ":")
		if len(values) != 2 {
			slog.Error(fmt.Sprintf("Invalid session cookie: %+v", values))
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}

		sessionID := values[0]
		userId := values[1]

		rCtx := r.Context()
		s, u, err := m.sessionService.GetWithUser(rCtx, sessionID, userId)
		if err != nil {
			slog.Error(err.Error())
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}

		ctx := context.WithValue(rCtx, LoggedUserKey, u)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
	return http.HandlerFunc(fn)
}

func (m *AuthMiddleware) RedirectIfLogged(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		sessionCookie, _ := r.Cookie(m.sessionCookieName)

		if sessionCookie == nil {
			next.ServeHTTP(w, r)
			return
		}
		if sessionCookie.Value == "" {
			next.ServeHTTP(w, r)
			return
		}

		rCtx := r.Context()
		s, _ := m.sessionService.GetByToken(rCtx, sessionCookie.Value)
		if s == nil {
			next.ServeHTTP(w, r)
			return
		}

		http.Redirect(w, r, "/home", http.StatusFound)
	}
	return http.HandlerFunc(fn)
}
