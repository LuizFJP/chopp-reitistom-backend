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
}
