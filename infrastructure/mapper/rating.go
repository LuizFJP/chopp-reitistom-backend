package mapper

import (
	"chopp-reitistom-backend/domain/entity"
	"chopp-reitistom-backend/infrastructure/persistence/model"
)

func FromRatingModelToEntity(model *model.Rating) *entity.Rating {
	return &entity.Rating{
		UUID:    model.UUID,
		Comment: model.Comment,
		Grade:   model.Grade,
	}
}

func FromRatingEntityToModel(entity *entity.Rating) *model.Rating {
	return &model.Rating{
		UUID:    entity.UUID,
		Comment: entity.Comment,
		Grade:   entity.Grade,
	}
}

func FromRatingModelToEntityArray(models []*model.Rating) []*entity.Rating {
	entityRatings := make([]*entity.Rating, len(models))
	for _, modelRating := range models {
		entityRatings = append(entityRatings, FromRatingModelToEntity(modelRating))
	}

	return entityRatings
}
