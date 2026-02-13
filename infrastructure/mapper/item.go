package mapper

import "chopp-reitistom-backend/infrastructure/persistence/model"

func CreateItemsModelFromQuantity(quantity, productId uint) []*model.Item {
	items := make([]*model.Item, 0)
	var i uint
	for ; i < quantity; i++ {
		item := &model.Item{ProductId: productId}
		items = append(items, item)
	}
	return items
}
