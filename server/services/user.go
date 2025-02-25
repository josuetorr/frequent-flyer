package services

import (
	"context"

	"github.com/josuetorr/frequent-flyer/server/models"
)

type UserService struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) Insert(ctx context.Context, user *models.User) error {
	return s.repo.Insert(ctx, user)
}
