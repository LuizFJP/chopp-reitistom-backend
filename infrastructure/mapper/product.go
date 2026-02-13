package mapper

import (
	"chopp-reitistom-backend/domain/entity"
	"chopp-reitistom-backend/infrastructure/persistence/model"
)

func FromProductModelToEntity(model *model.Product) *entity.Product {
	return &entity.Product{
		UUID:        model.UUID,
		Name:        model.Name,
		Price:       model.Price,
		Ratings:     FromRatingModelToEntityArray(model.Ratings),
		Ingredients: FromIngredientModelToEntityArray(model.Ingredients),
	}
}

func FromProductEntityToModel(entity *entity.Product) *model.Product {
	return &model.Product{
		UUID: entity.UUID,
		Name: entity.Name,
	}
}

func UpdateProductFromEntityToModel(entity *entity.Product, model *model.Product) {
	if entity.Name != model.Name {
		model.Name = entity.Name
	}
}

func FromProductFromModelToEntityArray(models []*model.Product) []*entity.Product {
	entities := make([]*entity.Product, 0)
	for _, m := range models {
		entities = append(entities, FromProductModelToEntity(m))
	}
	return entities
}
