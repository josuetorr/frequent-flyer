package services

import (
	"context"
	"errors"
	"log/slog"

	"github.com/josuetorr/frequent-flyer/internal/utils"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	userRepo UserRepository
}

func NewAuthService(userRepo UserRepository) *AuthService {
	return &AuthService{userRepo: userRepo}
}

func (s *AuthService) Signup(ctx context.Context, email string, password string) (ID, error) {
	u, _ := s.userRepo.GetByEmail(ctx, email)
	if u != nil {
		return "", errors.New("User already exists")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	user := &User{
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

	return "", nil
}

func (s *AuthService) Login(ctx context.Context, email string, password string) (ID, error) {
	u, err := s.userRepo.GetByEmail(ctx, email)
	if err != nil {
		slog.Error("Could not find user with given email: " + email)
		return "", errors.New("Invalid credentials")
	}

	if err := utils.ComparePassword(u.Password, password); err != nil {
		slog.Error("Invalid password: " + password)
		return "", errors.New("Invalid credentials")
	}

	return "", nil
}
