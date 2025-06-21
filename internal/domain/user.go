package domain

import "github.com/google/uuid"

type User struct {
	ID       uuid.UUID
	Email    string
	Nickname string
	PassHash string
}

func NewUser(email string, nickname string, passHash string) *User {
	return &User{
		ID:       uuid.New(),
		Email:    email,
		Nickname: nickname,
		PassHash: passHash,
	}
}
