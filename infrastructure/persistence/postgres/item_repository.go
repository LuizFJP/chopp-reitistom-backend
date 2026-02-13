package postgres

import (
	"chopp-reitistom-backend/infrastructure/persistence/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ItemRepository struct {
	db *gorm.DB
}

func (ir *ItemRepository) CreateBatch(items []*model.Item, quantity uint) (uint, error) {
	result := ir.db.CreateInBatches(items, int(quantity))
	if result.Error != nil {
		return 0, result.Error
	}
	return uint(result.RowsAffected), nil
}

func (ir *ItemRepository) Create(item *model.Item) error {
	result := ir.db.Create(item)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (ir *ItemRepository) Update(item *model.Item) error {
	result := ir.db.Save(item)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (ir *ItemRepository) GetByUUID(uuid uuid.UUID) (*model.Item, error) {
	var result model.Item
	err := ir.db.Find(&result, "uuid = ?", uuid).Error
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (ir *ItemRepository) Delete(item *model.Item) error {
	result := ir.db.Delete(item)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (ir *ItemRepository) GetQuantity(productId uint) (int, error) {
	var resultQuery []*model.Item
	err := ir.db.Find(&resultQuery, "productId = ?", productId).Error
	if err != nil {
		return 0, err
	}
	
	return len(resultQuery), nil
}
