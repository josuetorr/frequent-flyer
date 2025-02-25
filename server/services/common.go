package services

import (
	"context"

	"github.com/josuetorr/frequent-flyer/server/models"
)

type (
	ID   = models.ID
	User = models.User
)

type UserRepository interface {
	Insert(context.Context, *User) error
	Get(context.Context, ID) (*User, error)
}
