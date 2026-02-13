package repository

import (
	"chopp-reitistom-backend/infrastructure/persistence/model"

	"github.com/google/uuid"
)

type CategoryRepositoryInterface interface {
	Create(category *model.Category) error
	Update(category *model.Category) error
	GetByUUID(uuid uuid.UUID) (*model.Category, error)
	Delete(category *model.Category) error
}
