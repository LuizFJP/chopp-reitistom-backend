package postgres

import (
	"chopp-reitistom-backend/infrastructure/persistence/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SuggestionRepository struct {
	db *gorm.DB
}

func NewSuggestionRepository(
	db *gorm.DB) *SuggestionRepository {
	return &SuggestionRepository{db}
}

func (sr *SuggestionRepository) Create(model *model.Suggestion) error {
	result := sr.db.Create(model)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
func (sr *SuggestionRepository) Delete(model *model.Suggestion) error {
	result := sr.db.Delete(model)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (sr *SuggestionRepository) Update(model *model.Suggestion) error {
	result := sr.db.Save(model)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (sr *SuggestionRepository) GetAllByUser(userId uint) ([]*model.Suggestion, error) {
	var results []*model.Suggestion
	err := sr.db.Find(&results, "user_id = ?", userId).Error
	if err != nil {
		return nil, err
	}
	return results, nil
}

func (sr *SuggestionRepository) GetAll() ([]*model.Suggestion, error) {
	var results []*model.Suggestion
	err := sr.db.Find(&results).Error
	if err != nil {
		return nil, err
	}
	return results, nil
}

func (sr *SuggestionRepository) GetByUUID(uuid uuid.UUID) (*model.Suggestion, error) {
	var result *model.Suggestion
	err := sr.db.Find(&result, "uuid = ?", uuid).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}
