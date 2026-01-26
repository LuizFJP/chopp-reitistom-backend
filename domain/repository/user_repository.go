package repository

import (
	"chopp-reitistom-backend/infrastructure/persistence/model"

	"github.com/google/uuid"
)

type UserRepositoryInterface interface {
	Save(user *model.User) (int64, error)
	GetByEmail(email string) (*model.User, error)
	GetByUUID(uuid uuid.UUID) (*model.User, error)
	Update(user *model.User) (int64, error)
	Delete(user *model.User) (int64, error)
}
