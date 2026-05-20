package domain

import "github.com/google/uuid"

type User struct {
	ID       string
	Username string
	Password string
	Email    string
}

func NewUser(id, username, password, email string) (*User, error) {
	return &User{
		ID:       uuid.UUID,
		Username: username,
		Password: password,
		Email:    email,
	}, nil
}``
