package controller

import (
	"context"
	"fmt"

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
	// FIXME:
	fmt.Println("mutationResolver#Instruction")
	fmt.Printf("%#v\n", input)
	return &controllermodel.Instruction{}, nil
}
func (r *mutationResolver) Order(ctx context.Context, input controllermodel.NewOrder) (*controllermodel.Order, error) {
	// FIXME:
	fmt.Println("mutationResolver#Order")
	fmt.Printf("%#v\n", input)
	return &controllermodel.Order{}, nil
}
func (r *mutationResolver) Organization(ctx context.Context, input controllermodel.NewOrganization) (*controllermodel.Organization, error) {
	// FIXME:
	fmt.Println("mutationResolver#Organization")
	fmt.Printf("%#v\n", input)
	return &controllermodel.Organization{}, nil
}
func (r *mutationResolver) User(ctx context.Context, input controllermodel.NewUser) (*controllermodel.User, error) {
	// FIXME:
	fmt.Println("mutationResolver#User")
	fmt.Printf("%#v\n", input)
	return &controllermodel.User{}, nil
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Node(ctx context.Context, id string) (controllermodel.Node, error) {
	panic("not implemented")
}
func (r *queryResolver) Instruction(ctx context.Context, id string) (*controllermodel.Instruction, error) {
	// FIXME:
	fmt.Println("queryResolver#Instruction")
	fmt.Printf("id: %#v\n", id)
	return &controllermodel.Instruction{}, nil
}
func (r *queryResolver) Instructions(ctx context.Context) ([]controllermodel.Instruction, error) {
	// FIXME:
	fmt.Println("queryResolver#Instructions")
	return []controllermodel.Instruction{}, nil
}
func (r *queryResolver) Order(ctx context.Context, id string) (*controllermodel.Order, error) {
	// FIXME:
	fmt.Println("queryResolver#Order")
	return &controllermodel.Order{}, nil
}
func (r *queryResolver) Orders(ctx context.Context) ([]controllermodel.Order, error) {
	// FIXME:
	fmt.Println("queryResolver#Orders")
	return []controllermodel.Order{}, nil
}
func (r *queryResolver) Organization(ctx context.Context, id string) (*controllermodel.Organization, error) {
	// FIXME:
	fmt.Println("queryResolver#Organization")
	fmt.Printf("id: %#v\n", id)
	return &controllermodel.Organization{}, nil
}
func (r *queryResolver) Organizations(ctx context.Context) ([]controllermodel.Organization, error) {
	// FIXME:
	fmt.Println("queryResolver#Organizations")
	return []controllermodel.Organization{}, nil
}
func (r *queryResolver) User(ctx context.Context, id string) (*controllermodel.User, error) {
	// FIXME:
	fmt.Println("queryResolver#User")
	fmt.Printf("id: %#v\n", id)
	return &controllermodel.User{}, nil
}
func (r *queryResolver) Users(ctx context.Context) ([]controllermodel.User, error) {
	// FIXME:
	fmt.Println("queryResolver#Users")
	return []controllermodel.User{}, nil
}
