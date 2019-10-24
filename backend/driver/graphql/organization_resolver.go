package graphql

import (
	"context"

	"github.com/sky0621/fiktivt-handelssystem/adapter/controller"
)

func (r *mutationResolver) Organization(ctx context.Context, input controller.NewOrganization) (*controller.Organization, error) {
	panic("not implemented")
}
func (r *mutationResolver) User(ctx context.Context, input controller.NewUser) (*controller.User, error) {
	panic("not implemented")
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
