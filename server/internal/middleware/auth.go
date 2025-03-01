package middleware

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/josuetorr/frequent-flyer/server/handlers"
)

type LoggedUser = string

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

		rCtx := r.Context()
		s, err := m.sessionService.GetByToken(rCtx, sessionCookie.Value)
		if err != nil {
			slog.Error(err.Error())
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}

		ctx := context.WithValue(rCtx, LoggedUser(s.UserID), s.UserID)
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
