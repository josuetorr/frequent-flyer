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

func (s *SessionService) GetWithUser(ctx context.Context, sessionID models.ID, userID models.ID) (*models.Session, *models.User, error) {
	return s.repo.GetWithUser(ctx, sessionID, userID)
}
