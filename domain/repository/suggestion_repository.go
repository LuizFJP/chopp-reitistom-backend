package repository

import (
	"chopp-reitistom-backend/infrastructure/persistence/model"

	"github.com/google/uuid"
)

type SuggestionRepositoryInterface interface {
	Create(model *model.Suggestion) error
	Delete(model *model.Suggestion) error
	Update(model *model.Suggestion) error
	GetAllByUser(userId uint) ([]*model.Suggestion, error)
	GetAll() ([]*model.Suggestion, error)
	GetByUUID(uuid uuid.UUID) (*model.Suggestion, error)
}
