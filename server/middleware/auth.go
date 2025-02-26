package middleware

import (
	"errors"
	"log/slog"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/josuetorr/frequent-flyer/server/utils"
)

func AuthMiddleware(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			return []byte(utils.GetJWTSecret()), nil
		})

		// TODO: add the authed user to request context
		switch {
		case token.Valid:
			next.ServeHTTP(w, r)
			return
		case errors.Is(err, jwt.ErrTokenExpired):
			// TODO: refresh toke;
			slog.Error("Expired token: " + err.Error())
		case errors.Is(err, jwt.ErrTokenNotValidYet):
			slog.Error("Token not yet valid" + err.Error())
		default:
			slog.Error("Invalid token: " + err.Error())
		}
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
	}
	return http.HandlerFunc(fn)
}
