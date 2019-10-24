package graphql

import (
	"context"

	"github.com/sky0621/fiktivt-handelssystem/adapter/controller"
)

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct {
	adapter controller.GraphQLAdapter
}

func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Node(ctx context.Context, id string) (controller.Node, error) {
	panic("not implemented")
}

func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) Noop(ctx context.Context, input *controller.NoopInput) (*controller.NoopPayload, error) {
	panic("not implemented")
}
