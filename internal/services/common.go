package services

import (
	"context"

	"github.com/josuetorr/frequent-flyer/internal/models"
)

type (
	UserRepository interface {
		Insert(context.Context, *models.User) error
		GetById(context.Context, models.ID) (*models.User, error)
		GetByEmail(context.Context, string) (*models.User, error)
		Update(context.Context, models.ID, *models.User) error
		Delete(context.Context, models.ID, bool) error
	}
	SessionRepository interface {
		Insert(context.Context, *models.Session) error
		GetByToken(context.Context, models.SessionToken) (*models.Session, error)
		Update(context.Context, models.SessionToken, *models.Session) error
		Delete(context.Context, models.SessionToken) error
	}
)
