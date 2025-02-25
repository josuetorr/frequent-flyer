package handlers

import (
	"context"

	"github.com/josuetorr/frequent-flyer/server/models"
)

type (
	User = models.User
	ID   = models.ID
)

type UserService interface {
	Insert(context.Context, *models.User) error
	Get(context.Context, ID) (*User, error)
}
