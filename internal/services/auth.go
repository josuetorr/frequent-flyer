package services

import (
	"context"
	"errors"
	"log/slog"
	"time"

	"github.com/josuetorr/frequent-flyer/internal/models"
	"github.com/josuetorr/frequent-flyer/internal/utils"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	userRepo    UserRepository
	sessionRepo SessionRepository
}

func NewAuthService(userRepo UserRepository, sessionRepo SessionRepository) *AuthService {
	return &AuthService{userRepo: userRepo, sessionRepo: sessionRepo}
}

func (s *AuthService) Signup(ctx context.Context, email string, password string) (models.ID, error) {
	u, _ := s.userRepo.GetByEmail(ctx, email)
	if u != nil {
		return "", errors.New("User already exists")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	user := &models.User{
		Email:    email,
		Password: string(hash),
	}

	if err := s.userRepo.Insert(ctx, user); err != nil {
		return "", err
	}

	user, err = s.userRepo.GetByEmail(ctx, user.Email)
	if err != nil {
		return "", err
	}

	return user.ID, nil
}

func (s *AuthService) Login(ctx context.Context, email string, password string) (*models.Session, error) {
	u, err := s.userRepo.GetByEmail(ctx, email)
	if err != nil {
		slog.Error("Could not find user with given email: " + email)
		return nil, errors.New("Invalid credentials")
	}

	if err := utils.ComparePassword(u.Password, password); err != nil {
		slog.Error("Invalid password: " + password)
		return nil, errors.New("Invalid credentials")
	}

	token, err := utils.GenerateRandomToken()
	if err != nil {
		return nil, err
	}

	weekDuration := time.Hour * 24 * 7
	session := &models.Session{
		UserID:    u.ID,
		Token:     token,
		CreatedAt: time.Now().UTC(),
		ExpiresAt: time.Now().UTC().Add(weekDuration),
	}

	if err := s.sessionRepo.Insert(ctx, session); err != nil {
		return nil, err
	}

	return session, nil
}
