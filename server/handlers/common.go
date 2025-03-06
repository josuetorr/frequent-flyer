package handlers

import (
	"context"

	"github.com/josuetorr/frequent-flyer/internal/models"
)

type UserService interface {
	VerifyUser(ctx context.Context, userID models.ID) error
	GetByEmail(ctx context.Context, email string) (*models.User, error)
}

type AuthService interface {
	Signup(ctx context.Context, email string, password string) (models.ID, error)
	Login(ctx context.Context, email string, password string) (*models.Session, error)
}

type MailService interface {
	GenerateEmailLink(userID models.ID, endpoint string, secret string) string
	SendVerificationEmail(ctx context.Context, link string, to string) error
	SendPasswordResetEmail(ctx context.Context, link string, to string) error
}

type SessionService interface {
	GetWithUser(ctx context.Context, sessionID models.ID, userID models.ID) (*models.Session, *models.User, error)
}
