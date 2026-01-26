package mapper

import (
	"chopp-reitistom-backend/domain/entity"
	"chopp-reitistom-backend/infrastructure/persistence/model"
)

func FromAddressModelToEntity(model *model.Address) *entity.Address {
	return &entity.Address{
		UUID:         model.UUID,
		Street:       model.Street,
		Neighborhood: model.Neighborhood,
		Number:       model.Number,
		City:         model.City,
		Complement:   model.Complement,
		LandMark:     model.LandMark,
	}
}

func FromAddressEntityToModel(entity *entity.Address) *model.Address {
	return &model.Address{
		UUID:         entity.UUID,
		Street:       entity.Street,
		Neighborhood: entity.Neighborhood,
		Number:       entity.Number,
		City:         entity.City,
		Complement:   entity.Complement,
		LandMark:     entity.LandMark,
	}
}

func UpdateAddressFromEntityToModel(entity *entity.Address, model *model.Address) {
	if entity.LandMark != model.LandMark {
		model.LandMark = entity.LandMark
	}

	if entity.Complement != model.Complement {
		model.Complement = entity.Complement
	}

	if entity.Street != model.Street {
		model.Street = entity.Street
	}

	if entity.Number != model.Number {
		model.Number = entity.Number
	}

	if entity.Neighborhood != model.Neighborhood {
		model.Neighborhood = entity.Neighborhood
	}
}
