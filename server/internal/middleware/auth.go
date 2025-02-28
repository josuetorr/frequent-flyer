package middleware

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/josuetorr/frequent-flyer/internal/models"
	"github.com/josuetorr/frequent-flyer/server/handlers"
	"github.com/josuetorr/frequent-flyer/server/internal/utils"
)

func AuthMiddlerware(sessionService handlers.SessionService) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			sessionToken, err := r.Cookie(utils.SessionCookieName)
			if err != nil {
				slog.Error(err.Error())
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			s, err := sessionService.GetByToken(r.Context(), sessionToken.Value)
			if err != nil {
				slog.Error(err.Error())
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			ctx := context.WithValue(r.Context(), models.ID(s.UserID), s.UserID)
			next.ServeHTTP(w, r.WithContext(ctx))
		}
		return http.HandlerFunc(fn)
	}
}
