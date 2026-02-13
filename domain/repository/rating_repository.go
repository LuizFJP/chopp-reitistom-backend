package repository

import (
	"chopp-reitistom-backend/infrastructure/persistence/model"

	"github.com/google/uuid"
)

type RatingRepositoryInterface interface {
	Create(rating *model.Rating) error
	Delete(rating *model.Rating) error
	GetAllByProduct(productId uint) ([]*model.Rating, error)
	GetAll() ([]*model.Rating, error)
	GetByUUID(uuid uuid.UUID) (*model.Rating, error)
}
