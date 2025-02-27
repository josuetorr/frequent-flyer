package data

import (
	"context"
	"errors"
)

const (
	createUserQuery         = "INSERT INTO users (firstname, lastname, email, deleted_at, verified, password, refresh_token) VALUES ($1, $2, $3, $4, $5, $6, $7)"
	selectUserByIdQuery     = "SELECT * FROM users WHERE id = $1"
	selectUserByEmailQuery  = "SELECT * FROM users WHERE email = $1"
	selectRefreshTokenQuery = "SELECT refresh_token FROM users WHERE id = $1"
	updateUserQuery         = `
  UPDATE users
  SET firstname = $1, lastname = $2, email = $3, verified = $4
  WHERE id = $5
  `
	updateUserRefreshToken = `
  UPDATE users
  SET refresh_token = $1
  WHERE id =$2
  `
	deleteHardUserQuery = "DELETE FROM users WHERE id = $1"
	deleteSoftUserQuery = `
  UPDATE users
  SET deleted_at = NOW()
  WHERE id = $1
  `
)

type UserRepository struct {
	db *DBPool
}

func NewUserRepository(db *DBPool) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Insert(ctx context.Context, u *User) error {
	_, err := r.db.Exec(
		ctx,
		createUserQuery,
		u.Firstname,
		u.Lastname,
		u.Email,
		u.DeletedAt,
		u.Verified,
		u.Password,
		u.RefreshToken,
	)
	return err
}

func (r *UserRepository) get(ctx context.Context, by string, value string) (*User, error) {
	allowedQueries := map[string]string{
		"id":    selectUserByIdQuery,
		"email": selectUserByEmailQuery,
	}
	q, ok := allowedQueries[by]
	if !ok {
		return nil, errors.New("Invalid query field")
	}
	row := r.db.QueryRow(ctx, q, value)
	var u User
	err := row.Scan(&u.ID, &u.Firstname, &u.Lastname, &u.Email, &u.Verified, &u.DeletedAt, &u.Password, &u.RefreshToken)
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func (r *UserRepository) GetById(ctx context.Context, id ID) (*User, error) {
	return r.get(ctx, "id", id)
}

func (r *UserRepository) GetByEmail(ctx context.Context, email ID) (*User, error) {
	return r.get(ctx, "email", email)
}

func (r *UserRepository) Update(ctx context.Context, id ID, u *User) error {
	_, err := r.db.Exec(ctx, updateUserQuery, u.Firstname, u.Lastname, u.Email, u.Verified, id)
	return err
}

func (r *UserRepository) Delete(ctx context.Context, id ID, hard bool) error {
	var query string
	if hard {
		query = deleteHardUserQuery
	} else {
		query = deleteSoftUserQuery
	}
	_, err := r.db.Exec(ctx, query, id)
	return err
}

func (r *UserRepository) UpdateRefreshToken(ctx context.Context, id ID, refreshToken string) error {
	_, err := r.db.Exec(ctx, updateUserRefreshToken, refreshToken, id)
	return err
}

func (r *UserRepository) GetRefreshToken(ctx context.Context, id ID) (string, error) {
	row := r.db.QueryRow(ctx, selectRefreshTokenQuery, id)
	var refreshToken string
	if err := row.Scan(&refreshToken); err != nil {
		return "", err
	}
	return refreshToken, nil
}
