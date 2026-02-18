package postgres

import (
	"chopp-reitistom-backend/infrastructure/persistence/model"

	"gorm.io/gorm"
)

type ItemRepository struct {
	db *gorm.DB
}

func NewItemRepository(
	db *gorm.DB,
) *ItemRepository {
	return &ItemRepository{db}
}

func (ir *ItemRepository) CreateBatch(items []*model.Item, quantity uint) (uint, error) {
	result := ir.db.CreateInBatches(items, int(quantity))
	if result.Error != nil {
		return 0, result.Error
	}
	return uint(result.RowsAffected), nil
}

func (ir *ItemRepository) GetQuantity(productId uint) (int, error) {
	var resultQuery []*model.Item
	err := ir.db.Find(&resultQuery, "productId = ?", productId).Error
	if err != nil {
		return 0, err
	}

	return len(resultQuery), nil
}
