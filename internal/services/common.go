package services

import (
	"context"

	"github.com/josuetorr/frequent-flyer/internal/data"
	"github.com/josuetorr/frequent-flyer/server/models"
)

type (
	ID             = models.ID
	User           = models.User
	UserRepository interface {
		data.Repository[User]
		GetByEmail(context.Context, string) (*User, error)
		UpdateRefreshToken(ctx context.Context, id string, refreshToken string) error
		GetRefreshToken(ctx context.Context, id ID) (string, error)
	}
)
