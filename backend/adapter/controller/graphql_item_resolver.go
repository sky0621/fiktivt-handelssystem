package controller

import (
	"context"
)

func (r *Resolver) Item() ItemResolver {
	return &ItemResolverImpl{r: r}
}

type ItemResolverImpl struct {
	r *Resolver
}

func (i *ItemResolverImpl) ItemHolder(ctx context.Context, obj *Item) (*ItemHolder, error) {
	domainItemHolder, err := i.r.itemHolder.GetItemHolderByItemID(ctx, obj.ID)
	if err != nil {
		return nil, err
	}
	return ToControllerItemHolder(domainItemHolder), nil
}

func (r *Resolver) ItemHolder() ItemHolderResolver {
	return &ItemHolderResolverImpl{}
}

type ItemHolderResolverImpl struct {
	r *Resolver
}

func (i *ItemHolderResolverImpl) HoldItems(ctx context.Context, obj *ItemHolder) ([]Item, error) {
	domainItems, err := i.r.item.GetItemsByItemHolderID(ctx, obj.ID)
	if err != nil {
		return nil, err
	}
	var items []Item
	for _, domainItem := range domainItems {
		items = append(items, *ToControllerItem(domainItem))
	}
	return items, nil
}

/********************************************************************
 * Query
 */

func (r *queryResolver) Item(ctx context.Context, id string) (*Item, error) {
	domainItem, err := r.item.GetItem(ctx, id)
	if err != nil {
		return nil, err
	}
	return ToControllerItem(domainItem), nil
}

func (r *queryResolver) Items(ctx context.Context) ([]Item, error) {
	domainItems, err := r.item.GetItems(ctx)
	if err != nil {
		return nil, err
	}
	var items []Item
	for _, domainItem := range domainItems {
		items = append(items, *ToControllerItem(domainItem))
	}
	return items, nil
}

func (r *queryResolver) ItemHolder(ctx context.Context, id string) (*ItemHolder, error) {
	res, err := r.itemHolder.GetItemHolder(ctx, id)
	if err != nil {
		return nil, err
	}
	return ToControllerItemHolder(res), nil
}

func (r *queryResolver) ItemHolders(ctx context.Context) ([]ItemHolder, error) {
	results, err := r.itemHolder.GetItemHolders(ctx)
	if err != nil {
		return nil, err
	}
	var itemHolders []ItemHolder
	for _, res := range results {
		itemHolders = append(itemHolders, *ToControllerItemHolder(res))
	}
	return itemHolders, nil
}

/********************************************************************
 * Mutation
 */

func (r *mutationResolver) CreateItem(ctx context.Context, input ItemInput) (string, error) {
	return r.item.CreateItem(ctx, ToCommandItemModel(input))
}

func (r *mutationResolver) CreateItemHolder(ctx context.Context, input ItemHolderInput) (string, error) {
	return r.itemHolder.CreateItemHolder(ctx, ToCommandItemHolderModel(input))
}
