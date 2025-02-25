package data

import (
	"context"
)

const (
	createUserQuery = "INSERT INTO users (firstname, lastname, email, verified) VALUES ($1, $2, $3, $4)"
	selectUserQuery = "SELECT * FROM users WHERE id = $1"
	updateUserQuery = `
  UPDATE users
  SET firstname = $1, lastname = $2, email = $3, verified = $4
  WHERE id = $5
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

func NewUserRepositor(db *DBPool) Repository[User] {
	return &UserRepository{db: db}
}

func (r *UserRepository) Insert(ctx context.Context, u *User) error {
	_, err := r.db.Exec(ctx, createUserQuery, u.Firstname, u.Lastname, u.Email, u.Verified)
	return err
}

func (r *UserRepository) Get(ctx context.Context, id ID) (*User, error) {
	row := r.db.QueryRow(ctx, selectUserQuery, id)
	var u User
	err := row.Scan(&u.ID, &u.Firstname, &u.Lastname, &u.Email, &u.Verified)
	if err != nil {
		return nil, err
	}
	return &u, nil
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
