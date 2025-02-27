package api

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/josuetorr/frequent-flyer/internal/utils"
	"github.com/josuetorr/frequent-flyer/server/models"
)

type (
	User = models.User
	ID   = models.ID
)

type UserService interface {
	Insert(context.Context, *models.User) error
	GetById(context.Context, ID) (*User, error)
	Update(context.Context, ID, *User) error
	Delete(context.Context, ID, bool) error
	GetRefreshToken(context.Context, ID) (string, error)
}

type AuthService interface {
	Signup(ctx context.Context, email string, password string) (string, error)
	Login(ctx context.Context, email string, password string) (string, string, error)
}

type ApiHandleFn func(w http.ResponseWriter, r *http.Request) (*utils.ApiResponse, *utils.ApiError)

// ServeHTTP can take params a dependencies, such as other loggers
func (fn ApiHandleFn) ServeHTTP() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res, err := fn(w, r)
		if err != nil {
			// using default logger
			slog.Error(err.Error.Error())
			utils.WriteError(w, err)
			return
		}

		if res != nil {
			utils.WriteJSON(w, res.Status, res.Data)
		} else {
			w.WriteHeader(http.StatusOK)
		}
	}
}
