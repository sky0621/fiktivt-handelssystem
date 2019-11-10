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
		sortCondition *domain.SortConditionModel,
		searchDirectionType domain.SearchDirection,
		limit int, startCursor *string, endCursor *string,
	) ([]*domain.QueryItemHolderModel, int, error)
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
	sortCondition *domain.SortConditionModel,
	searchDirectionType domain.SearchDirection,
	limit int, startCursor *string, endCursor *string,
) ([]*domain.QueryItemHolderModel, int, error) {
	itemHolders, err := i.itemHolderDomain.GetItemHoldersByCondition(ctx,
		searchWordCondition, itemHolderCondition, sortCondition, searchDirectionType,
		limit, startCursor, endCursor)
	if err != nil {
		return nil, 0, err
	}
	allCount, err := i.itemHolderDomain.GetItemHolderAllCount(ctx)
	if err != nil {
		return nil, 0, err
	}
	return itemHolders, allCount, nil
}

func (i *itemHolder) CreateItemHolder(ctx context.Context, input domain.CommandItemHolderModel) (string, error) {
	return i.itemHolderDomain.CreateItemHolder(ctx, input)
}
