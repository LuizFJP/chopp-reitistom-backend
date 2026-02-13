package postgres

import (
	"chopp-reitistom-backend/infrastructure/persistence/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type RatingRepository struct {
	db *gorm.DB
}

func NewRatingRepository(
	db *gorm.DB,
) *RatingRepository {
	return &RatingRepository{db}
}

func (rr *RatingRepository) Create(rating *model.Rating) error {
	if err := rr.db.Create(rating).Error; err != nil {
		return err
	}
	return nil
}
func (rr *RatingRepository) Delete(rating *model.Rating) error {
	if err := rr.db.Delete(rating).Error; err != nil {
		return err
	}
	return nil
}

func (rr *RatingRepository) GetAllByProduct(productId uint) ([]*model.Rating, error) {
	var results []*model.Rating
	if err := rr.db.Find(results, "product_id = ?", productId).Error; err != nil {
		return nil, err
	}
	return results, nil
}

func (rr *RatingRepository) GetAll() ([]*model.Rating, error) {
	var results []*model.Rating
	if err := rr.db.Find(results).Error; err != nil {
		return nil, err
	}
	return results, nil
}

func (rr *RatingRepository) GetByUUID(uuid uuid.UUID) (*model.Rating, error) {
	var result model.Rating
	if err := rr.db.Find(&result, "uuid = ?", uuid).Error; err != nil {
		return nil, err
	}
	return &result, nil
}
