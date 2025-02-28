package data

import (
	"context"

	"github.com/josuetorr/frequent-flyer/internal/models"
)

const (
	createSessionQuery = `
  INSERT INTO sessions (user_id, token, user_agent, ip_address, created_at, expires_at) 
  VALUES ($1, $2, $3, $4, $5, $6)`
	selectSessionByTokenQuery = "SELECT * FROM sessions WHERE token = $1"
	deleteSessionByTokenQuery = "DELETE FROM sessions WHERE token = $1"
)

type SessionRepository struct {
	db *DBPool
}

func NewSessionRepository(db *DBPool) *SessionRepository {
	return &SessionRepository{db: db}
}

func (r *SessionRepository) Insert(ctx context.Context, s *models.Session) error {
	_, err := r.db.Query(ctx, createSessionQuery, s.UserID, s.Token, s.UserAgent, s.IpAddr, s.CreatedAt, s.ExpiresAt)
	return err
}

func (r *SessionRepository) GetByToken(ctx context.Context, token models.SessionToken) (*models.Session, error) {
	row, err := r.db.Query(ctx, selectSessionByTokenQuery, token)
	if err != nil {
		return nil, err
	}

	var s models.Session
	err = row.Scan(&s.ID, &s.UserID, &s.Token, &s.UserAgent, &s.IpAddr, &s.CreatedAt, &s.ExpiresAt)
	if err != nil {
		return nil, err
	}

	return &s, nil
}

func (r *SessionRepository) Update(ctx context.Context, token models.SessionToken, session *models.Session) error {
	panic("session repo update not implemented")
}

func (r *SessionRepository) Delete(ctx context.Context, token models.SessionToken) error {
	_, err := r.db.Exec(ctx, deleteSessionByTokenQuery, token)
	return err
}
