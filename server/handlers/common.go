package handlers

import (
	"context"

	"github.com/josuetorr/frequent-flyer/internal/models"
)

type AuthService interface {
	Signup(ctx context.Context, email string, password string) (models.ID, error)
	Login(ctx context.Context, email string, password string) (*models.Session, error)
}

type MailService interface {
	SendVerificationEmail(ctx context.Context, to string) error
}

type SessionService interface {
	GetWithUser(ctx context.Context, sessionID models.ID, userID models.ID) (*models.Session, *models.User, error)
}
