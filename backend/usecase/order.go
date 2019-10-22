package usecase

import "github.com/sky0621/fiktivt-handelssystem/domain"

func NewOrder(orderLogic domain.Order, organizationLogic domain.Organization) Order {
	return &order{
		orderLogic:        orderLogic,
		organizationLogic: organizationLogic,
	}
}

type Order interface {
	GetOrders() ([]*domain.OrderModel, error)
}

type order struct {
	orderLogic        domain.Order
	organizationLogic domain.Organization
}

func (o *order) GetOrders() ([]*domain.OrderModel, error) {
	// FIXME:
	return o.orderLogic.GetOrders()
}
