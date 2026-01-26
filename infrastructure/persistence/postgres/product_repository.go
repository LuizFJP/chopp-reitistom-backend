package postgres

import (
	"chopp-reitistom-backend/infrastructure/persistence/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (pr *ProductRepository) Create(product *model.Product) error              { return nil }
func (pr *ProductRepository) Update(product *model.Product) error              { return nil }
func (pr *ProductRepository) GetByUUID(uuid uuid.UUID) (*model.Product, error) { return nil, nil }
func (pr *ProductRepository) Delete(product *model.Product) error              { return nil }
