package services

import (
	"context"

	"github.com/josuetorr/frequent-flyer/internal/models"
)

type SessionService struct {
	repo SessionRepository
}

func NewSessionService(repo SessionRepository) *SessionService {
	return &SessionService{repo: repo}
}

func (s *SessionService) Insert(ctx context.Context, session *models.Session) error {
	return s.repo.Insert(ctx, session)
}

func (s *SessionService) GetById(ctx context.Context, id models.ID) (*models.Session, error) {
	return s.repo.GetById(ctx, id)
}
