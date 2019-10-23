package controller

import "github.com/sky0621/fiktivt-handelssystem/domain"

// Order
//   |- OrderDetail
//         |- Instruction

func FromOrderModels(models []*domain.OrderModel) []Order {
	var orders []Order
	for _, model := range models {
		orders = append(orders, FromOrderModel(model))
	}
	return orders
}

func FromOrderModel(model *domain.OrderModel) Order {
	user := FromUserModel(model.User)
	return Order{
		ID:           model.ID,
		Name:         model.Name,
		User:         &user,
		OrderDetails: nil, // FIXME:
	}
}

func FromOrderDetailModel(model *domain.OrderDetailModel) OrderDetail {
	return OrderDetail{
		ID:          model.ID,
		Order:       nil, // FIXME: 循環参照！
		Instruction: nil,
	}
}

func FromUserModel(model *domain.UserModel) User {
	return User{
		ID:            model.ID,
		Name:          model.Name,
		Organizations: nil,
	}
}
