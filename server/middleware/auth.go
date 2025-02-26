package middleware

import (
	"errors"
	"fmt"
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
		// TODO: clean this up
		switch {
		case token.Valid:
			next.ServeHTTP(w, r)
		case errors.Is(err, jwt.ErrTokenMalformed):
			http.Error(w, fmt.Sprint("That's not even a token"), http.StatusUnauthorized)
			return
		case errors.Is(err, jwt.ErrTokenSignatureInvalid):
			// Invalid signature
			http.Error(w, fmt.Sprint("Invalid signature"), http.StatusUnauthorized)
			return
		case errors.Is(err, jwt.ErrTokenExpired) || errors.Is(err, jwt.ErrTokenNotValidYet):
			// Token is either expired or not active yet
			http.Error(w, fmt.Sprint("Timing is everything"), http.StatusUnauthorized)
			return
		default:
			http.Error(w, fmt.Sprint("Couldn't handle this token:"), http.StatusUnauthorized)
			return
		}
	}
	return http.HandlerFunc(fn)
}
