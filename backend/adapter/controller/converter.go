package controller

import "github.com/sky0621/fiktivt-handelssystem/domain"

func FromOrderModels(models []*domain.OrderModel) []Order {
	var orders []Order
	for _, model := range models {
		orders = append(orders, FromOrderModel(model))
	}
	return orders
}

func FromOrderModel(model *domain.OrderModel) Order {
	return Order{
		ID:           model.ID,
		Name:         model.Name,
		User:         nil, // FIXME:
		OrderDetails: nil, // FIXME:
	}
}
