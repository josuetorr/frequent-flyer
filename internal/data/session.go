package data

import (
	"context"

	"github.com/josuetorr/frequent-flyer/internal/models"
)

const (
	createSessionQuery = `
  INSERT INTO sessions (user_id, user_agent, ip_address, created_at, expires_at) 
  VALUES ($1, $2, $3, $4, $5)
  RETURNING *`
	// TODO: this is where we are at. Fetch the session with the user with a join query
	selectSessionWithUserQuery = `
  SELECT 
    sessions.id,
    sessions.user_id,
    sessions.user_agent,
    sessions.ip_address,
    sessions.create_at,
    sessions.expired_at,
    users.id,
    users.firstname,
    users.lastname,
    users.email,
    users.deleted_at,
    users.verified
  FROM sessions 
  JOIN users ON session.user_id = user.id
  WHERE id = $1 AND user_id = $2
  `
	deleteSessionByTokenQuery = "DELETE FROM sessions WHERE token = $1 AND user_id = $2"
)

type SessionRepository struct {
	db *DBPool
}

func NewSessionRepository(db *DBPool) *SessionRepository {
	return &SessionRepository{db: db}
}

func (r *SessionRepository) Insert(ctx context.Context, s *models.Session) (*models.Session, error) {
	var session models.Session
	err := r.db.QueryRow(
		ctx,
		createSessionQuery,
		s.UserID,
		s.UserAgent,
		s.IpAddr,
		s.CreatedAt,
		s.ExpiresAt).
		Scan(
			&session.ID,
			&session.UserID,
			&session.UserAgent,
			&session.IpAddr,
			&session.CreatedAt,
			&session.ExpiresAt,
		)
	return &session, err
}

func (r *SessionRepository) GetWithUser(ctx context.Context, sessionID models.ID, userID models.ID) (*models.Session, *models.User, error) {
	row := r.db.QueryRow(ctx, selectSessionWithUserQuery, sessionID, userID)

	var s models.Session
	var u models.User
	err := row.Scan(
		&s.ID,
		&s.UserID,
		&s.UserAgent,
		&s.IpAddr,
		&s.CreatedAt,
		&s.ExpiresAt,
		&u.ID,
		&u.Firstname,
		&u.Lastname,
		&u.Email,
		&u.DeletedAt,
		&u.Verified,
	)
	if err != nil {
		return nil, nil, err
	}

	return &s, &u, nil
}

func (r *SessionRepository) Delete(ctx context.Context, sessionID models.ID, userID models.ID) error {
	_, err := r.db.Exec(ctx, deleteSessionByTokenQuery, sessionID, userID)
	return err
}
