package mapper

import (
	"chopp-reitistom-backend/domain/entity"
	"chopp-reitistom-backend/infrastructure/persistence/model"
)

func FromSuggestionModelToEntity(model *model.Suggestion) *entity.Suggestion {
	return &entity.Suggestion{
		UUID:        model.UUID,
		Description: model.Description,
	}
}

func FromSuggestionEntityToModel(entity *entity.Suggestion) *model.Suggestion {
	return &model.Suggestion{
		UUID:        entity.UUID,
		Description: entity.Description,
	}
}

func UpdateSuggestionFromEntityToModel(entity *entity.Suggestion, model *model.Suggestion) {
	if entity.Description != model.Description {
		model.Description = entity.Description
	}
}

func FromSuggestionModelToEntityArray(models []*model.Suggestion) []*entity.Suggestion {
	entitySuggestions := make([]*entity.Suggestion, len(models))
	for _, modelSuggestion := range models {
		entitySuggestions = append(entitySuggestions, FromSuggestionModelToEntity(modelSuggestion))
	}

	return entitySuggestions
}
