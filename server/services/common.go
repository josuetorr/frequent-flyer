package services

import (
	"context"

	"github.com/josuetorr/frequent-flyer/server/data"
	"github.com/josuetorr/frequent-flyer/server/models"
)

type (
	ID             = models.ID
	User           = models.User
	UserRepository interface {
		data.Repository[User]
		GetByEmail(context.Context, string) (*User, error)
	}
)
