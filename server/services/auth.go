package services

import (
	"context"
	"errors"

	"github.com/josuetorr/frequent-flyer/server/utils"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	userRepo UserRepository
}

func NewAuthService(userRepo UserRepository) *AuthService {
	return &AuthService{userRepo: userRepo}
}

func (s *AuthService) Signup(ctx context.Context, email string, password string) (string, error) {
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

	token := utils.NewJwtToken(user.ID)
	signedToken, err := utils.SignToken(token)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func (s *AuthService) Login(ctx context.Context, email string, password string) (string, string, error) {
	panic("implment auth service login")
}
