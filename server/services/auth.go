package services

import (
	"context"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/josuetorr/frequent-flyer/server/utils"
	"golang.org/x/crypto/bcrypt"
)

type SignupRequest struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type AuthService struct {
	userRepo UserRepository
}

func NewAuthService(userRepo UserRepository) *AuthService {
	return &AuthService{userRepo: userRepo}
}

func (s *AuthService) Signup(ctx context.Context, req *SignupRequest) (string, error) {
	u, _ := s.userRepo.GetByEmail(ctx, req.Email)
	if u != nil {
		return "", errors.New("User already exists")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	user := &User{
		Firstname: req.Firstname,
		Lastname:  req.Lastname,
		Email:     req.Lastname,
		Password:  string(hash),
	}

	if err := s.userRepo.Insert(ctx, user); err != nil {
		return "", err
	}

	user, err = s.userRepo.GetByEmail(ctx, user.Email)
	if err != nil {
		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss": "frequent-flyer",
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	signedToken, err := token.SignedString([]byte(utils.GetJWTSecret()))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}
