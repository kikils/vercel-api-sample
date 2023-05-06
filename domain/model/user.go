package model

import "github.com/google/uuid"

type User struct {
	ID   string
	Name string
}

func NewUser(name string) (*User, error) {
	uid, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	return &User{
		ID:   uid.String(),
		Name: name,
	}, nil
}
