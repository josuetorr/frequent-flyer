package handlers

import (
	"context"

	"github.com/josuetorr/frequent-flyer/internal/models"
)

type AuthService interface {
	Signup(ctx context.Context, email string, password string) (models.ID, error)
	Login(ctx context.Context, email string, password string) (*models.Session, error)
	Logout(ctx context.Context, token models.SessionToken) error
}

type SessionService interface {
	GetByToken(ctx context.Context, token models.SessionToken) (*models.Session, error)
}
