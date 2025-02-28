package data

import (
	"context"

	"github.com/josuetorr/frequent-flyer/internal/models"
)

type (
	ID      = models.ID
	User    = models.User
	Session = models.Session
)

type Repository[T any] interface {
	Insert(context.Context, *T) error
	GetById(context.Context, ID) (*T, error)
	Update(context.Context, ID, *T) error
	Delete(context.Context, ID, bool) error
}
