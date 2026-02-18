package repository

import (
	"chopp-reitistom-backend/infrastructure/persistence/model"
)

type ItemRepositoryInterface interface {
	CreateBatch(items []*model.Item, quantity uint) (uint, error)
	GetQuantity(productId uint) (int, error)
}
