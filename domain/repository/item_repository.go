package repository

import (
	"chopp-reitistom-backend/infrastructure/persistence/model"

	"github.com/google/uuid"
)

type ItemRepositoryInterface interface {
	Create(item *model.Item) error
	CreateBatch(items []*model.Item, quantity uint) (uint, error)
	Update(item *model.Item) error
	GetByUUID(uuid uuid.UUID) (*model.Item, error)
	Delete(item *model.Item) error
	GetQuantity(productId uint) (int, error)
}
