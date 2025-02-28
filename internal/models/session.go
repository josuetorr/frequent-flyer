package models

import "time"

type Session struct {
	ID        ID
	UserID    ID
	Token     string
	UserAgent string
	IpAddr    string
	CreatedAt time.Time
	ExpiresIn time.Duration
}
