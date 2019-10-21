package controller

import (
	"context"

	controllermodel "github.com/sky0621/fiktivt-handelssystem/adapter/controller/model"
	"github.com/sky0621/fiktivt-handelssystem/driver"
)

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct{}

func (r *Resolver) Mutation() driver.MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() driver.QueryResolver {
	return &queryResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) Noop(ctx context.Context, input *controllermodel.NoopInput) (*controllermodel.NoopPayload, error) {
	panic("not implemented")
}
func (r *mutationResolver) Instruction(ctx context.Context, input controllermodel.NewInstruction) (*controllermodel.Instruction, error) {
	panic("not implemented")
}
func (r *mutationResolver) Order(ctx context.Context, input controllermodel.NewOrder) (*controllermodel.Order, error) {
	panic("not implemented")
}
func (r *mutationResolver) Organization(ctx context.Context, input controllermodel.NewOrganization) (*controllermodel.Organization, error) {
	panic("not implemented")
}
func (r *mutationResolver) User(ctx context.Context, input controllermodel.NewUser) (*controllermodel.User, error) {
	panic("not implemented")
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Node(ctx context.Context, id string) (controllermodel.Node, error) {
	panic("not implemented")
}
func (r *queryResolver) Instruction(ctx context.Context, id string) (*controllermodel.Instruction, error) {
	panic("not implemented")
}
func (r *queryResolver) Instructions(ctx context.Context) ([]controllermodel.Instruction, error) {
	panic("not implemented")
}
func (r *queryResolver) Order(ctx context.Context, id string) (*controllermodel.Order, error) {
	panic("not implemented")
}
func (r *queryResolver) Orders(ctx context.Context) ([]controllermodel.Order, error) {
	panic("not implemented")
}
func (r *queryResolver) Organization(ctx context.Context, id string) (*controllermodel.Organization, error) {
	panic("not implemented")
}
func (r *queryResolver) Organizations(ctx context.Context) ([]controllermodel.Organization, error) {
	panic("not implemented")
}
func (r *queryResolver) User(ctx context.Context, id string) (*controllermodel.User, error) {
	panic("not implemented")
}
func (r *queryResolver) Users(ctx context.Context) ([]controllermodel.User, error) {
	panic("not implemented")
}
