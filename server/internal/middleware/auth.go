package middleware

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/josuetorr/frequent-flyer/server/handlers"
	"github.com/josuetorr/frequent-flyer/server/internal/utils"
)

type LoggedUser = string

type Middleware = func(http.Handler) http.Handler

func AuthMiddlerware(sessionService handlers.SessionService) Middleware {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			sessionCookie, err := r.Cookie(utils.SessionCookieName)
			if err != nil {
				slog.Error(err.Error())
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			rCtx := r.Context()
			s, err := sessionService.GetByToken(rCtx, sessionCookie.Value)
			if err != nil {
				slog.Error(err.Error())
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			ctx := context.WithValue(rCtx, LoggedUser(s.UserID), s.UserID)
			next.ServeHTTP(w, r.WithContext(ctx))
		}
		return http.HandlerFunc(fn)
	}
}

func RedirectIfLogged(sessionService handlers.SessionService) Middleware {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			sessionCookie, _ := r.Cookie(utils.SessionCookieName)

			if sessionCookie == nil {
				next.ServeHTTP(w, r)
				return
			}
			if sessionCookie.Value == "" {
				next.ServeHTTP(w, r)
				return
			}

			rCtx := r.Context()
			s, _ := sessionService.GetByToken(rCtx, sessionCookie.Value)
			if s == nil {
				next.ServeHTTP(w, r)
				return
			}

			w.Header().Set("HX-REDIRECT", "/home")
		}
		return http.HandlerFunc(fn)
	}
}
