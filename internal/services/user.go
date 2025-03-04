package services

import (
	"context"

	"github.com/josuetorr/frequent-flyer/internal/models"
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

func (s *UserService) GetById(ctx context.Context, id models.ID) (*models.User, error) {
	return s.repo.GetById(ctx, id)
}

func (s *UserService) GetByEmail(ctx context.Context, id models.ID) (*models.User, error) {
	return s.repo.GetByEmail(ctx, id)
}

func (s *UserService) Update(ctx context.Context, id models.ID, u *models.User) error {
	return s.repo.Update(ctx, id, u)
}

func (s *UserService) Delete(ctx context.Context, id models.ID, hard bool) error {
	return s.repo.Delete(ctx, id, hard)
}

func (s *UserService) VerifyUser(ctx context.Context, userID models.ID) error {
	u, err := s.GetById(ctx, userID)
	if err != nil {
		return err
	}

	u.Verified = true
	if err := s.Update(ctx, userID, u); err != nil {
		return err
	}

	return nil
}
