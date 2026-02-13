package entity

import "github.com/google/uuid"

type Product struct {
	UUID         uuid.UUID     `json:"UUID"`
	Name         string        `json:"name"`
	Price        float64       `json:"price"`
	CategoryUUID uuid.UUID     `json:"categoryUUID"`
	Ratings      []*Rating     `json:"ratings"`
	Ingredients  []*Ingredient `json:"ingredients"`
	Quantity     uint          `json:"quantity"`
}
