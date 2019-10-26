package controller

import (
	"context"

	"github.com/sky0621/fiktivt-handelssystem/usecase"
)

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

func NewResolverRoot(item usecase.Item, itemHolder usecase.ItemHolder) ResolverRoot {
	return &Resolver{
		item:       item,
		itemHolder: itemHolder,
	}
}

type Resolver struct {
	item       usecase.Item
	itemHolder usecase.ItemHolder
}

func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Node(ctx context.Context, id string) (Node, error) {
	panic("not implemented")
}

func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) Noop(ctx context.Context, input *NoopInput) (*NoopPayload, error) {
	panic("not implemented")
}
