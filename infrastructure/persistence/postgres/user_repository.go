package postgres

import (
	"chopp-reitistom-backend/infrastructure/persistence/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (ur *UserRepository) Save(user *model.User) (int64, error) {
	result := ur.db.Create(user)
	if err := result.Error; err != nil {
		return 0, err
	}
	return result.RowsAffected, nil
}

func (ur *UserRepository) GetByEmail(email string) (*model.User, error) {
	var result model.User
	err := ur.db.Find(&result, "email = ?", email).Error
	if err != nil {
		return nil, err
	}

	resultEmpty := result.Id == 0

	if resultEmpty {
		return nil, nil
	}

	return &result, nil
}

func (ur *UserRepository) GetByUUID(uuid uuid.UUID) (*model.User, error) {
	var result model.User
	err := ur.db.Find(&result, "uuid = ?", uuid).Error
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (ur *UserRepository) Update(user *model.User) (int64, error) {
	result := ur.db.Save(user)
	if err := result.Error; err != nil {
		return 0, err
	}
	return result.RowsAffected, nil
}

func (ur *UserRepository) Delete(user *model.User) (int64, error) {
	result := ur.db.Delete(user)
	if err := result.Error; err != nil {
		return 0, err
	}
	return result.RowsAffected, nil
}
