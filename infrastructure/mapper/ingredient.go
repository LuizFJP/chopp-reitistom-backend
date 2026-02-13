package mapper

import (
	"chopp-reitistom-backend/domain/entity"
	"chopp-reitistom-backend/infrastructure/persistence/model"

	"github.com/google/uuid"
)

func FromIngredientModelToEntity(model *model.Ingredient) *entity.Ingredient {
	return &entity.Ingredient{
		UUID: model.UUID,
		Name: model.Name,
	}
}

func FromIngredientEntityToModel(entity *entity.Ingredient) *model.Ingredient {
	return &model.Ingredient{
		UUID: entity.UUID,
		Name: entity.Name,
	}
}

func FromIngredientModelToEntityArray(models []*model.Ingredient) []*entity.Ingredient {
	entityIngredients := make([]*entity.Ingredient, len(models))
	for _, modelIngredient := range models {
		entityIngredients = append(entityIngredients, FromIngredientModelToEntity(modelIngredient))
	}

	return entityIngredients
}

func UpdateIngredientFromEntityToModel(entity *entity.Ingredient, model *model.Ingredient) {
	if entity.Name != model.Name {
		model.Name = entity.Name
	}
}

func ExtractIngredientUUIDs(ingredients []*entity.Ingredient) []uuid.UUID {
	results := make([]uuid.UUID, 0)
	for _, ingredient := range ingredients {
		results = append(results, ingredient.ProductId)
	}
	return results
}
