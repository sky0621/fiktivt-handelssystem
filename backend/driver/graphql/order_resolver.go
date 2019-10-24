package graphql

import (
	"context"

	"github.com/sky0621/fiktivt-handelssystem/adapter/controller"
)

func (r *queryResolver) Order(ctx context.Context, id string) (*controller.Order, error) {
	panic("not implemented")
}

func (r *queryResolver) Orders(ctx context.Context) ([]controller.Order, error) {
	orders, err := r.adapter.GetOrderUsecase().GetOrders()
	if err != nil {
		// TODO: エラーハンドリング！
		return nil, err
	}
	return controller.FromOrderModels(orders), nil
}

func (r *queryResolver) Instruction(ctx context.Context, id string) (*controller.Instruction, error) {
	panic("not implemented")
}

func (r *queryResolver) Instructions(ctx context.Context) ([]controller.Instruction, error) {
	panic("not implemented")
}

func (r *mutationResolver) Order(ctx context.Context, input controller.NewOrder) (*controller.Order, error) {
	panic("not implemented")
}

func (r *mutationResolver) Instruction(ctx context.Context, input controller.NewInstruction) (*controller.Instruction, error) {
	panic("not implemented")
}
