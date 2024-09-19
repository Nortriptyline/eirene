package domain

import "github.com/google/uuid"

type User struct {
	ID          uuid.UUID
	ClientToken string
}

func NewUser(token string) *User {
	return &User{
		ID:          uuid.New(),
		ClientToken: token,
	}
}
