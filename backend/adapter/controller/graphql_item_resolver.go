package controller

import (
	"context"
	"github.com/sky0621/fiktivt-handelssystem/adapter/controller/converter"

	"github.com/sky0621/fiktivt-handelssystem/domain"
)

func (r *queryResolver) Item(ctx context.Context, id string) (*Item, error) {
	res, err := r.item.GetItem(id)
	if err != nil {
		return nil, err
	}
	return &Item{
		ID:    res.ID,
		Name:  res.Name,
		Price: res.Price,
		ItemHolder: &ItemHolder{
			ID:        res.ItemHolder.ID,
			Name:      res.ItemHolder.Name,
			Nickname:  res.ItemHolder.Nickname,
			HoldItems: res.ItemHolder.HoldItems,
		},
	}, nil
}

func (r *queryResolver) Items(ctx context.Context) ([]Item, error) {
	results, err := r.item.GetItems()
	if err != nil {
		return nil, err
	}
	var items []Item
	for _, res := range results {
		items = append(items, Item{
			ID:    res.ID,
			Name:  res.Name,
			Price: res.Price,
			ItemHolder: &ItemHolder{
				ID:        res.ItemHolder.ID,
				Name:      res.ItemHolder.Name,
				Nickname:  res.ItemHolder.Nickname,
				HoldItems: []Item{{
					ID:         res.ItemHolder.HoldItems.,
					Name:       "",
					Price:      0,
					ItemHolder: nil,
				}},
			},
		})
	}
	return items, nil
}

func (r *queryResolver) ItemHolder(ctx context.Context, id string) (*ItemHolder, error) {
	res, err := r.itemHolder.GetItemHolder(id)
	if err != nil {
		return nil, err
	}
	return converter.ToItemHolder(res), nil
}

func (r *queryResolver) ItemHolders(ctx context.Context) ([]ItemHolder, error) {
	results, err := r.itemHolder.GetItemHolders()
	if err != nil {
		return nil, err
	}
	var itemHolders []ItemHolder
	for _, res := range results {
		itemHolders = append(itemHolders, *converter.ToItemHolder(res))
	}
	return itemHolders, nil
}

func (r *mutationResolver) CreateItem(ctx context.Context, input ItemInput) (string, error) {
	return r.item.CreateItem(converter.ToCommandItemModel(input))
}

func (r *mutationResolver) CreateItemHolder(ctx context.Context, input ItemHolderInput) (string, error) {
	return r.itemHolder.CreateItemHolder(converter.ToCommandItemHolderModel(input))
}
