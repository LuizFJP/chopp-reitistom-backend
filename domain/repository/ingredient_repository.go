package repository

import (
	"chopp-reitistom-backend/infrastructure/persistence/model"

	"github.com/google/uuid"
)

type IngredientRepositoryInterface interface {
	Create(ingredient *model.Ingredient) error
	Delete(ingredient *model.Ingredient) error
	GetAll() ([]*model.Ingredient, error)
	GetByUUID(uuid uuid.UUID) (*model.Ingredient, error)
	Update(ingredient *model.Ingredient) error
	GetAllByProductId(productId uint) ([]*model.Ingredient, error)
	GetManyByUUID(uuids []uuid.UUID) ([]*model.Ingredient, error)
}
