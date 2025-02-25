package data

import (
	"context"
)

const (
	createUserQuery = "INSERT INTO users (firstname, lastname, email, verified) VALUES ($1, $2, $3, $4)"
	selectUserQuery = "SELECT * FROM users WHERE id = $1"
)

type UserRepository struct {
	db *DBPool
}

func NewUserRepositor(db *DBPool) *UserRepository {
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
