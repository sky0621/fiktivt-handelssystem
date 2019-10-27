package usecase

import (
	"context"

	"github.com/sky0621/fiktivt-handelssystem/domain"
)

func NewItem(itemDomain domain.Item) Item {
	return &item{
		itemDomain: itemDomain,
	}
}

type Item interface {
	GetItem(ctx context.Context, id string) (*domain.QueryItemModel, error)
	GetItems(ctx context.Context) ([]*domain.QueryItemModel, error)
	CreateItem(ctx context.Context, input domain.CommandItemModel) (string, error)
}

type item struct {
	itemDomain domain.Item
}

func (i *item) GetItem(ctx context.Context, id string) (*domain.QueryItemModel, error) {
	return i.itemDomain.GetItem(ctx, id)
}

func (i *item) GetItems(ctx context.Context) ([]*domain.QueryItemModel, error) {
	return i.itemDomain.GetItems(ctx)
}

func (i *item) CreateItem(ctx context.Context, input domain.CommandItemModel) (string, error) {
	return i.itemDomain.CreateItem(ctx, input)
}
