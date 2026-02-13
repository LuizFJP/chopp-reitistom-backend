package mapper

import (
	"chopp-reitistom-backend/domain/entity"
	"chopp-reitistom-backend/infrastructure/persistence/model"
)

func FromCategoryModelToEntity(model *model.Category) *entity.Category {
	return &entity.Category{
		UUID: model.UUID,
		Name: model.Name,
	}
}

func FromCategoryEntityToModel(entity *entity.Category) *model.Category {
	return &model.Category{
		UUID: entity.UUID,
		Name: entity.Name,
	}
}

func UpdateCategoryFromEntityToModel(entity *entity.Category, model *model.Category) {
	if entity.Name != model.Name {
		model.Name = entity.Name
	}
}
