package data

import "context"

const (
	createSessionQuery     = "INSERT INTO sessions (user_id, created_at, expires_in) VALUES ($1, $2, $3)"
	selectSessionByIdQuery = "SELECT * FROM sessions WHERE id = $1"
)

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

func (r *SessionRepository) GetById(ctx context.Context, id ID) (*Session, error) {
	row, err := r.db.Query(ctx, selectSessionByIdQuery, id)
	if err != nil {
		return nil, err
	}

	var s Session
	err = row.Scan(&s.ID, &s.UserID, &s.CreatedAt, &s.ExpiresIn)
	if err != nil {
		return nil, err
	}

	return &s, nil
}
