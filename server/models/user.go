package models

import "time"

type ID = string

type User struct {
	ID           ID         `json:"id"`
	Firstname    string     `json:"firstname"`
	Lastname     string     `json:"lastname"`
	Email        string     `json:"email"`
	DeletedAt    *time.Time `json:"deleted_at"`
	Verified     bool       `json:"verified"`
	Password     string     `json:"-"`
	RefreshToken string     `json:"refresh_token"`
}
