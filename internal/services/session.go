package services

import (
	"context"
)

type SessionService struct {
	repo SessionRepository
}

func NewSessionService(repo SessionRepository) *SessionService {
	return &SessionService{repo: repo}
}

func (s *SessionService) Insert(ctx context.Context, session *Session) error {
	return s.repo.Insert(ctx, session)
}

func (s *SessionService) GetById(ctx context.Context, id ID) (*Session, error) {
	return s.repo.GetById(ctx, id)
}
