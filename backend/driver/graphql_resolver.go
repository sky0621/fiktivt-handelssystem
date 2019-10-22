package driver

import (
	"context"

	"github.com/sky0621/fiktivt-handelssystem/adapter/controller"
)

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct {
	adapter controller.GraphQLAdapter
}

func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) Noop(ctx context.Context, input *controller.NoopInput) (*controller.NoopPayload, error) {
	panic("not implemented")
}
func (r *mutationResolver) Instruction(ctx context.Context, input controller.NewInstruction) (*controller.Instruction, error) {
	panic("not implemented")
}
func (r *mutationResolver) Order(ctx context.Context, input controller.NewOrder) (*controller.Order, error) {
	panic("not implemented")
}
func (r *mutationResolver) Organization(ctx context.Context, input controller.NewOrganization) (*controller.Organization, error) {
	panic("not implemented")
}
func (r *mutationResolver) User(ctx context.Context, input controller.NewUser) (*controller.User, error) {
	panic("not implemented")
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Node(ctx context.Context, id string) (controller.Node, error) {
	panic("not implemented")
}
func (r *queryResolver) Instruction(ctx context.Context, id string) (*controller.Instruction, error) {
	panic("not implemented")
}
func (r *queryResolver) Instructions(ctx context.Context) ([]controller.Instruction, error) {
	panic("not implemented")
}
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

func (r *queryResolver) Organization(ctx context.Context, id string) (*controller.Organization, error) {
	panic("not implemented")
}
func (r *queryResolver) Organizations(ctx context.Context) ([]controller.Organization, error) {
	panic("not implemented")
}
func (r *queryResolver) User(ctx context.Context, id string) (*controller.User, error) {
	panic("not implemented")
}
func (r *queryResolver) Users(ctx context.Context) ([]controller.User, error) {
	panic("not implemented")
}
