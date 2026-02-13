package entity

import "github.com/google/uuid"

type Category struct {
	UUID uuid.UUID `json:"uuid"`
	Name string    `json:"name"`
}
