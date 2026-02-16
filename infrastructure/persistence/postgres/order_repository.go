package postgres

import (
	"chopp-reitistom-backend/infrastructure/persistence/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type OrderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *OrderRepository {
	return &OrderRepository{db}
}

func (or *OrderRepository) Create(model *model.Order) error {
	result := or.db.Create(model)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (or *OrderRepository) Update(model *model.Order) error {
	result := or.db.Save(model)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (or *OrderRepository) GetByUUID(uuid uuid.UUID) (*model.Order, error) {
	var result model.Order
	err := or.db.Find(&result, "uuid = ?", uuid).Error
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (or *OrderRepository) GetManyByUserId(userId uint) ([]*model.Order, error) {
	var results []*model.Order
	err := or.db.Find(&results, "userId = ?", userId).Error
	if err != nil {
		return nil, err
	}
	return results, nil
}

func (or *OrderRepository) Delete(model *model.Order) error {
	result := or.db.Delete(model)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
