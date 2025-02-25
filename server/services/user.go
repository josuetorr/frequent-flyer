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

func (s *UserService) Get(ctx context.Context, id ID) (*User, error) {
	return s.repo.Get(ctx, id)
}

func (s *UserService) Update(ctx context.Context, id ID, u *User) error {
	return s.repo.Update(ctx, id, u)
}

func (s *UserService) Delete(ctx context.Context, id ID, hard bool) error {
	return s.repo.Delete(ctx, id, hard)
}
