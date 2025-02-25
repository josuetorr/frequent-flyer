package handlers

import (
	"context"

	"github.com/josuetorr/frequent-flyer/server/models"
)

type User = models.User

type UserService interface {
	Insert(context.Context, *models.User) error
}
