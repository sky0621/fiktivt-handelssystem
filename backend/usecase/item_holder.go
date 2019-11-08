package usecase

import (
	"context"

	"github.com/sky0621/fiktivt-handelssystem/domain"
)

func NewItemHolder(itemHolderDomain domain.ItemHolder) ItemHolder {
	return &itemHolder{itemHolderDomain: itemHolderDomain}
}

type ItemHolder interface {
	GetItemHolder(ctx context.Context, id string) (*domain.QueryItemHolderModel, error)
	GetItemHolders(ctx context.Context) ([]*domain.QueryItemHolderModel, error)
	GetItemHoldersByCondition(ctx context.Context, searchWord *string, first *int, after *string) (*domain.ItemHolderConnectionModel, error)
	CreateItemHolder(ctx context.Context, input domain.CommandItemHolderModel) (string, error)
}

type itemHolder struct {
	itemHolderDomain domain.ItemHolder
}

func (i *itemHolder) GetItemHolder(ctx context.Context, id string) (*domain.QueryItemHolderModel, error) {
	return i.itemHolderDomain.GetItemHolder(ctx, id)
}

func (i *itemHolder) GetItemHolders(ctx context.Context) ([]*domain.QueryItemHolderModel, error) {
	return i.itemHolderDomain.GetItemHolders(ctx)
}

func (i *itemHolder) GetItemHoldersByCondition(ctx context.Context, searchWord *string, first *int, after *string) (*domain.ItemHolderConnectionModel, error) {
	return i.itemHolderDomain.GetItemHoldersByCondition(ctx, searchWord, first, after)
}

func (i *itemHolder) CreateItemHolder(ctx context.Context, input domain.CommandItemHolderModel) (string, error) {
	return i.itemHolderDomain.CreateItemHolder(ctx, input)
}
