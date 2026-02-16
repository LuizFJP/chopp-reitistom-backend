package entity

import "github.com/google/uuid"

type Suggestion struct {
	UUID        uuid.UUID `json:"uuid"`
	Description string    `json:"description"`
}
