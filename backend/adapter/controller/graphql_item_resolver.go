package controller

import (
	"context"
	"fmt"

	"github.com/99designs/gqlgen/graphql"

	"github.com/sky0621/fiktivt-handelssystem/adapter/controller/model"
)

/********************************************************************
 * ItemResolver
 */

type itemResolver struct{ *Resolver }

func (r *Resolver) Item() ItemResolver {
	r.logger.Log("call")

	return &itemResolver{r}
}

func (r *itemResolver) ItemHolder(ctx context.Context, obj *model.Item) (*model.ItemHolder, error) {
	r.logger.Log("call")

	domainItemHolder, err := r.itemHolder.GetItemHolder(ctx, obj.ItemHolderID)
	if err != nil {
		r.logger.Log(err.Error())
		return nil, err
	}

	ret := ToControllerItemHolder(domainItemHolder)
	r.logger.Log(fmt.Sprintf("%#v", ret))
	return ret, nil
}

/********************************************************************
 * ItemHolderResolver
 */

type itemHolderResolver struct{ *Resolver }

func (r *Resolver) ItemHolder() ItemHolderResolver {
	r.logger.Log("call")

	return &itemHolderResolver{r}
}

func (r *itemHolderResolver) HoldItems(ctx context.Context, obj *model.ItemHolder) ([]model.Item, error) {
	r.logger.Log("call")

	domainItems, err := r.item.GetItemsByItemHolderID(ctx, obj.ID)
	if err != nil {
		r.logger.Log(err.Error())
		return nil, err
	}
	var items []model.Item
	for _, domainItem := range domainItems {
		items = append(items, *ToControllerItem(domainItem))
	}

	r.logger.Log(fmt.Sprintf("%#v", items))
	return items, nil
}

/********************************************************************
 * Query
 */

func (r *queryResolver) Item(ctx context.Context, id string) (*model.Item, error) {
	r.logger.Log("call")

	domainItem, err := r.item.GetItem(ctx, id, graphql.CollectAllFields(ctx))
	if err != nil {
		r.logger.Log(err.Error())
		return nil, err
	}

	ret := ToControllerItem(domainItem)
	r.logger.Log(fmt.Sprintf("%#v", ret))
	return ret, nil
}

func (r *queryResolver) Items(ctx context.Context) ([]model.Item, error) {
	r.logger.Log("call")

	domainItems, err := r.item.GetItems(ctx)
	if err != nil {
		r.logger.Log(err.Error())
		return nil, err
	}
	var items []model.Item
	for _, domainItem := range domainItems {
		items = append(items, *ToControllerItem(domainItem))
	}

	r.logger.Log(fmt.Sprintf("%#v", items))
	return items, nil
}

func (r *queryResolver) ItemHolder(ctx context.Context, id string) (*model.ItemHolder, error) {
	r.logger.Log("call")

	res, err := r.itemHolder.GetItemHolder(ctx, id)
	if err != nil {
		r.logger.Log(err.Error())
		return nil, err
	}

	ret := ToControllerItemHolder(res)
	r.logger.Log(fmt.Sprintf("%#v", ret))
	return ret, nil
}

func (r *queryResolver) ItemHolders(ctx context.Context) ([]model.ItemHolder, error) {
	r.logger.Log("call")

	results, err := r.itemHolder.GetItemHolders(ctx)
	if err != nil {
		r.logger.Log(err.Error())
		return nil, err
	}
	var itemHolders []model.ItemHolder
	for _, res := range results {
		itemHolders = append(itemHolders, *ToControllerItemHolder(res))
	}

	r.logger.Log(fmt.Sprintf("%#v", itemHolders))
	return itemHolders, nil
}

/********************************************************************
 * Mutation
 */

func (r *mutationResolver) CreateItem(ctx context.Context, input ItemInput) (string, error) {
	r.logger.Log("call")

	res, err := r.item.CreateItem(ctx, ToCommandItemModel(input))
	if err != nil {
		r.logger.Log(err.Error())
	}
	return res, nil
}

func (r *mutationResolver) CreateItemHolder(ctx context.Context, input ItemHolderInput) (string, error) {
	r.logger.Log("call")

	res, err := r.itemHolder.CreateItemHolder(ctx, ToCommandItemHolderModel(input))
	if err != nil {
		r.logger.Log(err.Error())
	}
	return res, nil
}
