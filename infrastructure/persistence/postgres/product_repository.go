package postgres

import (
	"chopp-reitistom-backend/infrastructure/persistence/model"
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (pr *ProductRepository) Create(product *model.Product) error {
	result := pr.db.Create(product)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (pr *ProductRepository) Update(product *model.Product) error {
	result := pr.db.Save(product)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (pr *ProductRepository) GetByUUID(uuid uuid.UUID) (*model.Product, error) {
	var result model.Product
	err := pr.db.Find(&result, "uuid = ?", uuid).Error
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (pr *ProductRepository) Delete(product *model.Product) error {
	result := pr.db.Delete(product)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (pr *ProductRepository) GetAll() ([]*model.Product, error) {
	var results []*model.Product
	err := pr.db.Find(&results).Error
	if err != nil {
		return nil, err
	}
	return results, nil
}

func (pr *ProductRepository) GetManyByUUID(uuids []uuid.UUID) ([]*model.Product, error) {
	var results []*model.Product
	err := pr.db.Find(&results, "uuid IN ?", uuids).Error
	if err != nil {
		return nil, err
	}
	return results, nil
}

func (pr *ProductRepository) AddIngredients(product *model.Product, ingredients []*model.Ingredient) error {
	return fmt.Errorf("not implemented yet")
}

func (pr *ProductRepository) RemoveIngredients(product *model.Product, ingredients []*model.Ingredient) error {
	return fmt.Errorf("not implemented yet")
}
