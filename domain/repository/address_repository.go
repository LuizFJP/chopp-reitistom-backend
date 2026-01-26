package repository

import (
	"chopp-reitistom-backend/infrastructure/persistence/model"

	"github.com/google/uuid"
)

type AddressRepositoryInterface interface {
	Create(address *model.Address) error
	Update(address *model.Address) error
	GetByUUID(uuid uuid.UUID) (*model.Address, error)
	Delete(address *model.Address) error
}
