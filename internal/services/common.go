package services

import (
	"context"

	"github.com/josuetorr/frequent-flyer/internal/data"
)

type (
	ID             = data.ID
	User           = data.User
	Session        = data.Session
	UserRepository interface {
		data.Repository[User]
		GetByEmail(context.Context, string) (*User, error)
		UpdateRefreshToken(ctx context.Context, id string, refreshToken string) error
		GetRefreshToken(ctx context.Context, id ID) (string, error)
	}
	SessionRepository interface {
		data.Repository[Session]
	}
)
