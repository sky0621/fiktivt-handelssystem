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
	GetItemHoldersByCondition(ctx context.Context,
		searchWordCondition *domain.SearchWordConditionModel,
		itemHolderCondition *domain.SearchItemHolderConditionModel,
		limit int, after *string,
		sortCondition *domain.SortConditionModel) ([]*domain.ItemHolder, int, error)
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

func (i *itemHolder) GetItemHoldersByCondition(ctx context.Context,
	searchWordCondition *domain.SearchWordConditionModel,
	itemHolderCondition *domain.SearchItemHolderConditionModel,
	limit int, after *string,
	sortCondition *domain.SortConditionModel) ([]*domain.ItemHolder, int, error) {
	return i.itemHolderDomain.GetItemHoldersByCondition(ctx, searchWordCondition, itemHolderCondition, limit, after, sortCondition)
}

func (i *itemHolder) CreateItemHolder(ctx context.Context, input domain.CommandItemHolderModel) (string, error) {
	return i.itemHolderDomain.CreateItemHolder(ctx, input)
}
