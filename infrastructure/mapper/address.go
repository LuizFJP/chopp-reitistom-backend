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
		Landmark:     model.Landmark,
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
		Landmark:     entity.Landmark,
	}
}

func UpdateAddressFromEntityToModel(entity *entity.Address, model *model.Address) {
	if entity.Landmark != model.Landmark {
		model.Landmark = entity.Landmark
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
