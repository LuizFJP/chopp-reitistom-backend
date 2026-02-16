package mapper

import (
	"chopp-reitistom-backend/domain/entity"
	"chopp-reitistom-backend/infrastructure/persistence/model"
)

func FromOrderModelToEntity(model *model.Order) (*entity.Order, error) {
	status, err := entity.ParseStatus(model.Status)
	if err != nil {
		return nil, err
	}

	return &entity.Order{
		UUID:          model.UUID,
		ProductsUUID:  ExtractUUIDsFromArray(model.Products),
		Status:        status,
		CreatedAt:     model.CreatedAt,
		ConfirmedAt:   model.ConfirmedAt,
		SentAt:        model.SentAt,
		CanceledAt:    model.CanceledAt,
		PaymentMethod: model.PaymentMethod,
		StatusPayment: model.StatusPayment,
	}, nil
}

func FromOrderEntityToModel(entity *entity.Order) *model.Order {
	return &model.Order{
		UUID:          entity.UUID,
		Status:        string(entity.Status),
		CreatedAt:     entity.CreatedAt,
		ConfirmedAt:   entity.ConfirmedAt,
		SentAt:        entity.SentAt,
		CanceledAt:    entity.CanceledAt,
		PaymentMethod: entity.PaymentMethod,
		StatusPayment: entity.StatusPayment,
	}
}

func FromOrderModelToEntityArray(models []*model.Order) ([]*entity.Order, error) {
	var entities []*entity.Order
	for _, m := range models {
		entityMapped, err := FromOrderModelToEntity(m)
		if err != nil {
			return nil, err
		}

		entities = append(entities, entityMapped)
	}

	return entities, nil
}
