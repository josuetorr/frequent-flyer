package models

type ID = string

type User struct {
	ID        ID     `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
	Verified  bool   `json:"verified"`
}
