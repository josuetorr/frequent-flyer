package services

import (
	"context"
	"errors"
	"log/slog"
	"time"

	"github.com/josuetorr/frequent-flyer/internal/models"
	"github.com/josuetorr/frequent-flyer/internal/utils"
)

var (
	InvalidCredentialError = errors.New("Invalid credentials")
	UserAlreadyExistsError = errors.New("User already exists")
)

type AuthService struct {
	userRepo    UserRepository
	sessionRepo SessionRepository
}

func NewAuthService(userRepo UserRepository, sessionRepo SessionRepository) *AuthService {
	return &AuthService{
		userRepo:    userRepo,
		sessionRepo: sessionRepo,
	}
}

// NOTE: we are assuming email is valid
func (s *AuthService) Signup(ctx context.Context, email string, password string) (models.ID, error) {
	u, _ := s.userRepo.GetByEmail(ctx, email)
	if u != nil {
		return "", UserAlreadyExistsError
	}

	hash, err := utils.HashPassword(password)
	if err != nil {
		return "", err
	}

	user := &models.User{
		Email:    email,
		Password: string(hash),
	}

	user, err = s.userRepo.Insert(ctx, user)
	if err != nil {
		return "", err
	}

	return user.ID, nil
}

func (s *AuthService) Login(ctx context.Context, email string, password string) (*models.Session, error) {
	u, err := s.userRepo.GetByEmail(ctx, email)
	if err != nil {
		slog.Error("Could not find user with given email: " + email)
		return nil, InvalidCredentialError
	}

	if err := utils.ComparePassword(u.Password, password); err != nil {
		slog.Error("Invalid password: " + password)
		return nil, InvalidCredentialError
	}

	weekDuration := time.Hour * 24 * 7
	session := &models.Session{
		UserID:    u.ID,
		CreatedAt: time.Now().UTC(),
		ExpiresAt: time.Now().UTC().Add(weekDuration),
	}

	session, err = s.sessionRepo.Insert(ctx, session)
	if err != nil {
		return nil, err
	}

	return session, nil
}
