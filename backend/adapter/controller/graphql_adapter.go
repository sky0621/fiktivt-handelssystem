package controller

import (
	"github.com/sky0621/fiktivt-handelssystem/usecase"
)

func NewGraphQLAdapter(orderUsecase usecase.Order, itemUsecase usecase.Item) GraphQLAdapter {
	return &graphQLAdapter{
		orderUsecase: orderUsecase,
		itemUsecase:  itemUsecase,
	}
}

type GraphQLAdapter interface {
	GetOrderUsecase() usecase.Order
	GetItemUsecase() usecase.Item
}

type graphQLAdapter struct {
	orderUsecase usecase.Order
	itemUsecase  usecase.Item
	// TODO: ユースケース追加！
}

func (a *graphQLAdapter) GetOrderUsecase() usecase.Order {
	return a.orderUsecase
}

func (a *graphQLAdapter) GetItemUsecase() usecase.Item {
	return a.itemUsecase
}
