package handlers

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/josuetorr/frequent-flyer/server/models"
	"github.com/josuetorr/frequent-flyer/server/services"
	"github.com/josuetorr/frequent-flyer/server/utils"
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
}

type AuthService interface {
	Signup(context.Context, *services.SignupRequest) (string, error)
}

type ApiHandleFn[T any] func(w http.ResponseWriter, r *http.Request) (*utils.ApiResponse[T], *utils.ApiError)

// ServeHTTP can take params a dependencies, such as other loggers
func (fn ApiHandleFn[T]) ServeHTTP() http.HandlerFunc {
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
