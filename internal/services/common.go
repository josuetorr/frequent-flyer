package services

import (
	"context"

	"github.com/josuetorr/frequent-flyer/internal/data"
	"github.com/josuetorr/frequent-flyer/internal/models"
)

type (
	UserRepository interface {
		data.Repository[models.User]
		GetByEmail(context.Context, string) (*models.User, error)
		UpdateRefreshToken(ctx context.Context, id string, refreshToken string) error
		GetRefreshToken(ctx context.Context, id models.ID) (string, error)
	}
	SessionRepository interface {
		data.Repository[models.Session]
		GetByUserId(ctx context.Context, userId models.ID) (*models.Session, error)
	}
)
