package entity

import "github.com/google/uuid"

type Suggestion struct {
	UUID        uuid.UUID `json:"UUID"`
	Description string    `json:"description"`
}
