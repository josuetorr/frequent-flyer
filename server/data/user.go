package data

import (
	"context"

	"github.com/josuetorr/frequent-flyer/server/models"
)

const createUserStmt = "INSERT INTO users (firstname,lastname,email) VALUES ($1,$2,$3)"

type UserRepository struct {
	db *DBPool
}

func NewUserRepositor(db *DBPool) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Insert(ctx context.Context, user *models.User) error {
	_, err := r.db.Exec(ctx, createUserStmt, user)
	return err
}
