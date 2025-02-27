package models

import "time"

type Session struct {
	ID        ID
	UserID    ID
	CreatedAt *time.Time
	ExpiresId uint
}
