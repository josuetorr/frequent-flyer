package services

import (
	"context"

	"github.com/josuetorr/frequent-flyer/server/models"
)

type UserRepository interface {
	Insert(context.Context, *models.User) error
}
