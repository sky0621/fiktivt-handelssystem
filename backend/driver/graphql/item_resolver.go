package graphql

import (
	"context"

	"github.com/sky0621/fiktivt-handelssystem/adapter/controller"
)

func (r *queryResolver) Item(ctx context.Context, id string) (*controller.Item, error) {
	r.adapter.GetItemUsecase()
	panic("implement me")
}

func (r *queryResolver) Items(ctx context.Context) ([]controller.Item, error) {
	panic("implement me")
}

func (r *queryResolver) ItemHolder(ctx context.Context, id string) (*controller.ItemHolder, error) {
	panic("implement me")
}

func (r *queryResolver) ItemHolders(ctx context.Context) ([]controller.ItemHolder, error) {
	panic("implement me")
}

func (r *mutationResolver) CreateItem(ctx context.Context, input controller.ItemInput) (string, error) {
	panic("implement me")
}

func (r *mutationResolver) CreateItemHolder(ctx context.Context, input controller.ItemHolderInput) (string, error) {
	panic("implement me")
}
