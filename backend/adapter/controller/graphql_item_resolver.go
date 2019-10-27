package controller

import (
	"context"
)

func (r *queryResolver) Item(ctx context.Context, id string) (*Item, error) {
	domainItem, err := r.item.GetItem(id)
	if err != nil {
		return nil, err
	}
	return ToControllerItem(domainItem), nil
}

func (r *queryResolver) Items(ctx context.Context) ([]Item, error) {
	domainItems, err := r.item.GetItems()
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
	res, err := r.itemHolder.GetItemHolder(id)
	if err != nil {
		return nil, err
	}
	return ToControllerItemHolder(res), nil
}

func (r *queryResolver) ItemHolders(ctx context.Context) ([]ItemHolder, error) {
	results, err := r.itemHolder.GetItemHolders()
	if err != nil {
		return nil, err
	}
	var itemHolders []ItemHolder
	for _, res := range results {
		itemHolders = append(itemHolders, *ToControllerItemHolder(res))
	}
	return itemHolders, nil
}

func (r *mutationResolver) CreateItem(ctx context.Context, input ItemInput) (string, error) {
	return r.item.CreateItem(ToCommandItemModel(input))
}

func (r *mutationResolver) CreateItemHolder(ctx context.Context, input ItemHolderInput) (string, error) {
	return r.itemHolder.CreateItemHolder(ToCommandItemHolderModel(input))
}
