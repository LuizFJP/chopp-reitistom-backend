package repository

import (
	"chopp-reitistom-backend/infrastructure/persistence/model"

	"github.com/google/uuid"
)

type OrderRepositoryInterface interface {
	Create(model *model.Order) error
	Update(model *model.Order) error
	GetByUUID(uuid uuid.UUID) (*model.Order, error)
	GetManyByUserId(userId uint) ([]*model.Order, error)
	Delete(model *model.Order) error
}
