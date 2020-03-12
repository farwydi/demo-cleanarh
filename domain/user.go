package domain

import "github.com/google/uuid"

type User struct {
	ID    ID
	Name  string
	Email string
}

func (u User) GetIDString() string {
	return uuid.UUID(u.ID).String()
}
