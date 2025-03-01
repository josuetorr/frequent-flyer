package data

import (
	"context"

	"github.com/josuetorr/frequent-flyer/internal/models"
)

const (
	createSessionQuery = `
  INSERT INTO sessions (user_id, user_agent, ip_address, created_at, expires_at) 
  VALUES ($1, $2, $3, $4, $5, $6)`
	selectSessionByTokenQuery = "SELECT * FROM sessions WHERE id = $1 AND user_id = $2"
	deleteSessionByTokenQuery = "DELETE FROM sessions WHERE token = $1 AND user_id = $2"
)

type SessionRepository struct {
	db *DBPool
}

func NewSessionRepository(db *DBPool) *SessionRepository {
	return &SessionRepository{db: db}
}

func (r *SessionRepository) Insert(ctx context.Context, s *models.Session) error {
	_, err := r.db.Query(ctx, createSessionQuery, s.UserID, s.UserAgent, s.IpAddr, s.CreatedAt, s.ExpiresAt)
	return err
}

func (r *SessionRepository) GetWithUser(ctx context.Context, sessionID models.ID, userID models.ID) (*models.Session, *models.User, error) {
	row := r.db.QueryRow(ctx, selectSessionByTokenQuery, sessionID, userID)

	var s models.Session
	err := row.Scan(&s.ID, &s.UserID, &s.UserAgent, &s.IpAddr, &s.CreatedAt, &s.ExpiresAt)
	if err != nil {
		return nil, nil, err
	}

	// TODO: also fetch user with JOIN
	panic("session.GetWithUser: need to figure out how to also fetch user")
	// return &s, nil, nil
}

func (r *SessionRepository) Delete(ctx context.Context, sessionID models.ID, userID models.ID) error {
	_, err := r.db.Exec(ctx, deleteSessionByTokenQuery, sessionID, userID)
	return err
}
