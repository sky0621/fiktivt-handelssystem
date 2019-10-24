package controller

import (
	"github.com/sky0621/fiktivt-handelssystem/usecase"
)

func NewGraphQLAdapter(orderUsecase usecase.Order) GraphQLAdapter {
	return &graphQLAdapter{
		orderUsecase: orderUsecase,
	}
}

type GraphQLAdapter interface {
	GetOrderUsecase() usecase.Order
}

type graphQLAdapter struct {
	orderUsecase usecase.Order
	// TODO: ユースケース追加！
}

func (a *graphQLAdapter) GetOrderUsecase() usecase.Order {
	return a.orderUsecase
}
