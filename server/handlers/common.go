package handlers

import (
	"context"

	"github.com/josuetorr/frequent-flyer/internal/models"
)

type AuthService interface {
	Signup(ctx context.Context, email string, password string) error
	Login(ctx context.Context, email string, password string) (models.Session, error)
}
