package entity

import "github.com/google/uuid"

type ProductIngredients struct {
	ProductUUID     uuid.UUID   `json:"productUUID"`
	IngredientsUUID []uuid.UUID `json:"ingredientsUUID"`
}

type Product struct {
	UUID         uuid.UUID     `json:"uuid"`
	Name         string        `json:"name"`
	Price        float64       `json:"price"`
	CategoryUUID uuid.UUID     `json:"categoryUUID"`
	Ratings      []*Rating     `json:"ratings"`
	Ingredients  []*Ingredient `json:"ingredients"`
	Quantity     uint          `json:"quantity"`
}
