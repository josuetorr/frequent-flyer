package models

import "time"

type SessionToken = string

type Session struct {
	ID        ID
	UserID    ID
	Token     SessionToken
	UserAgent string
	IpAddr    string
	CreatedAt time.Time
	ExpiresAt time.Time
}

func (s Session) Lifetime() int {
	return int(s.ExpiresAt.Sub(s.CreatedAt).Seconds())
}
