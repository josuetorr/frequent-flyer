package services

import (
	"context"

	"github.com/josuetorr/frequent-flyer/internal/models"
)

type (
	UserRepository interface {
		Insert(ctx context.Context, userID *models.User) error
		GetById(ctx context.Context, userID models.ID) (*models.User, error)
		GetByEmail(ctx context.Context, email string) (*models.User, error)
		Update(ctx context.Context, userID models.ID, user *models.User) error
		Delete(ctx context.Context, userID models.ID, hard bool) error
	}
	SessionRepository interface {
		Insert(ctx context.Context, session *models.Session) error
		GetWithUser(ctx context.Context, sessionID models.ID, userID models.ID) (*models.Session, *models.User, error)
		Delete(ctx context.Context, sessionID models.ID, userID models.ID) error
	}
)
