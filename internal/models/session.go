package models

import "time"

type Session struct {
	ID        ID
	UserID    ID
	UserAgent string
	IpAddr    string
	CreatedAt time.Time
	ExpiresAt time.Time
}

func (s Session) Lifetime() int {
	return int(s.ExpiresAt.Sub(s.CreatedAt).Seconds())
}

func (s Session) Expired() bool {
	return s.Lifetime() <= 0
}
