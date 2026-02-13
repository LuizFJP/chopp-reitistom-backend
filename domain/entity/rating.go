package entity

import "github.com/google/uuid"

type Rating struct {
	UUID      uuid.UUID `json:"uuid"`
	ProductId uuid.UUID `json:"productId"`
	Comment   string    `json:"comment"`
	Grade     uint      `json:"grade"`
}
