package repository

import (
	"chopp-reitistom-backend/infrastructure/persistence/model"

	"github.com/google/uuid"
)

type ProductRepositoryInterface interface {
	Create(product *model.Product) error
	Update(product *model.Product) error
	GetByUUID(uuid uuid.UUID) (*model.Product, error)
	Delete(product *model.Product) error
	GetAll() ([]*model.Product, error)
	AddIngredients(product *model.Product, ingredients []*model.Ingredient) error
	RemoveIngredients(product *model.Product, ingredients []*model.Ingredient) error
}
