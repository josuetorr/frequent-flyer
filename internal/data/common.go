package data

import (
	"context"

	"github.com/josuetorr/frequent-flyer/internal/models"
)

type Repository[T any] interface {
	Insert(context.Context, *T) error
	GetById(context.Context, models.ID) (*T, error)
	Update(context.Context, models.ID, *T) error
	Delete(context.Context, models.ID, bool) error
}
