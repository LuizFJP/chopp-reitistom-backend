package postgres

import (
	"chopp-reitistom-backend/infrastructure/persistence/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type IngredientRepository struct {
	db *gorm.DB
}

func NewIngredientRepository(
	db *gorm.DB,
) *IngredientRepository {
	return &IngredientRepository{db}
}

func (ir *IngredientRepository) Create(ingredient *model.Ingredient) error {
	result := ir.db.Create(ingredient)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (ir *IngredientRepository) Delete(ingredient *model.Ingredient) error {
	result := ir.db.Delete(ingredient)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (ir *IngredientRepository) GetAll() ([]*model.Ingredient, error) {
	var results []*model.Ingredient
	err := ir.db.Find(&results).Error
	if err != nil {
		return nil, err
	}
	return results, nil
}

func (ir *IngredientRepository) GetByUUID(uuid uuid.UUID) (*model.Ingredient, error) {
	var result model.Ingredient
	err := ir.db.Find(&result, "uuid = ?", uuid).Error
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (ir *IngredientRepository) GetManyByUUID(uuids []uuid.UUID) ([]*model.Ingredient, error) {
	var results []*model.Ingredient
	err := ir.db.Find(&results, "uuid IN ?", uuids).Error
	if err != nil {
		return nil, err
	}
	return results, nil
}

func (ir *IngredientRepository) GetAllByProductId(productId uint) ([]*model.Ingredient, error) {
	var results []*model.Ingredient
	err := ir.db.Find(&results, "product_id = ?", productId).Error
	if err != nil {
		return nil, err
	}
	return results, nil
}

func (ir *IngredientRepository) Update(ingredient *model.Ingredient) error {
	result := ir.db.Save(ingredient)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
