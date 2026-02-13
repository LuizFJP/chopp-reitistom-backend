package postgres

import (
	"chopp-reitistom-backend/infrastructure/persistence/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CategoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(
	db *gorm.DB,
) *CategoryRepository {
	return &CategoryRepository{db}
}

func (cr *CategoryRepository) Create(category *model.Category) error {
	result := cr.db.Create(category)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (cr *CategoryRepository) Update(category *model.Category) error {
	result := cr.db.Save(category)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (cr *CategoryRepository) GetByUUID(uuid uuid.UUID) (*model.Category, error) {
	var result model.Category
	err := cr.db.Find(&result, "uuid = ?", uuid).Error
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (cr *CategoryRepository) Delete(category *model.Category) error {
	result := cr.db.Delete(category)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
