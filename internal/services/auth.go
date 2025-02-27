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

func (s *AuthService) SignupWithJwt(ctx context.Context, email string, password string) (string, error) {
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

	signedToken, err := utils.NewAccessToken(user.ID)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func (s *AuthService) LoginWithJwt(ctx context.Context, email string, password string) (string, string, error) {
	u, err := s.userRepo.GetByEmail(ctx, email)
	if err != nil {
		slog.Error("Could not find user with given email: " + email)
		return "", "", errors.New("Invalid credentials")
	}

	if err := utils.ComparePassword(u.Password, password); err != nil {
		slog.Error("Invalid password: " + password)
		return "", "", errors.New("Invalid credentials")
	}

	accessToken, err := utils.NewAccessToken(u.ID)
	if err != nil {
		return "", "", err
	}

	refreshToken, err := utils.NewRefreshToken(u.ID)
	if err != nil {
		return "", "", err
	}

	if err := s.userRepo.UpdateRefreshToken(ctx, u.ID, refreshToken); err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}
