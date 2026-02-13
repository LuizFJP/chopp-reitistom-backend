package postgres

import (
	"chopp-reitistom-backend/infrastructure/persistence/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AddressRepository struct {
	db *gorm.DB
}

func NewAddressRepository(db *gorm.DB) *AddressRepository {
	return &AddressRepository{db: db}
}

func (aruc *AddressRepository) Create(address *model.Address) error {
	result := aruc.db.Create(address)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (aruc *AddressRepository) Update(address *model.Address) error {
	result := aruc.db.Save(address)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (aruc *AddressRepository) GetByUUID(uuid uuid.UUID) (*model.Address, error) {
	var result model.Address
	err := aruc.db.Find(&result, "uuid = ?", uuid).Error
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (aruc *AddressRepository) Delete(address *model.Address) error {
	result := aruc.db.Delete(address)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
