package middleware

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/josuetorr/frequent-flyer/server/handlers"
	"github.com/josuetorr/frequent-flyer/server/internal/utils"
)

type LoggedUser = string

func AuthMiddlerware(sessionService handlers.SessionService) func(http.Handler) http.Handler {
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
