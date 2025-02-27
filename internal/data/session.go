package data

import "context"

const createSessionQuery = "INSERT INTO sessions (user_id, created_at, expires_in) VALUES ($1, $2, $3)"

type SessionRepository struct {
	db *DBPool
}

func NewSessionRepository(db *DBPool) *SessionRepository {
	return &SessionRepository{db: db}
}

func (r *SessionRepository) Insert(ctx context.Context, s *Session) error {
	_, err := r.db.Exec(ctx, createSessionQuery, s.UserID, s.CreatedAt, s.ExpiresIn)
	return err
}
