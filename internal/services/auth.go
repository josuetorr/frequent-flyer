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

func (s *AuthService) Login(ctx context.Context, email string, password string) (models.ID, error) {
	u, err := s.userRepo.GetByEmail(ctx, email)
	if err != nil {
		slog.Error("Could not find user with given email: " + email)
		return "", errors.New("Invalid credentials")
	}

	if err := utils.ComparePassword(u.Password, password); err != nil {
		slog.Error("Invalid password: " + password)
		return "", errors.New("Invalid credentials")
	}

	weekDuration := time.Hour * 24 * 7
	session := &models.Session{
		UserID:    u.ID,
		CreatedAt: time.Now().UTC(),
		ExpiresIn: weekDuration,
	}

	if err := s.sessionRepo.Insert(ctx, session); err != nil {
		return "", err
	}

	session, err = s.sessionRepo.GetByUserId(ctx, u.ID)
	if err != nil {
		return "", err
	}

	return session.ID, nil
}
