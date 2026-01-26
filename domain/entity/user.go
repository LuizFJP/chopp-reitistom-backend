package entity

import "github.com/google/uuid"

type User struct {
	Email    string    `json:"email"`
	Name     string    `json:"name"`
	UUID     uuid.UUID `json:"id"`
	Password string    `json:"password"`
}

type SignIn struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
