package entity

import "github.com/google/uuid"

type Ingredient struct {
	UUID      uuid.UUID `json:"uuid"`
	ProductId uuid.UUID `json:"productId"`
	Name      string    `json:"name"`
}
