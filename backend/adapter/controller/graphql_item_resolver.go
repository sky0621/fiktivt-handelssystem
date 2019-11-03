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
	return &itemResolver{r}
}

func (r *itemResolver) ItemHolder(ctx context.Context, obj *model.Item) (*model.ItemHolder, error) {
	domainItemHolder, err := r.itemHolder.GetItemHolder(ctx, obj.ItemHolderID)
	if err != nil {
		return nil, err
	}
	return ToControllerItemHolder(domainItemHolder), nil
}

/********************************************************************
 * ItemHolderResolver
 */

type itemHolderResolver struct{ *Resolver }

func (r *Resolver) ItemHolder() ItemHolderResolver {
	return &itemHolderResolver{r}
}

func (r *itemHolderResolver) HoldItems(ctx context.Context, obj *model.ItemHolder) ([]model.Item, error) {
	rsctx := graphql.GetResolverContext(ctx)
	fmt.Println(rsctx)

	domainItems, err := r.item.GetItemsByItemHolderID(ctx, obj.ID)
	if err != nil {
		return nil, err
	}
	var items []model.Item
	for _, domainItem := range domainItems {
		items = append(items, *ToControllerItem(domainItem))
	}
	return items, nil
}

/********************************************************************
 * Query
 */

func (r *queryResolver) Item(ctx context.Context, id string) (*model.Item, error) {
	domainItem, err := r.item.GetItem(ctx, id, graphql.CollectAllFields(ctx))
	if err != nil {
		return nil, err
	}
	return ToControllerItem(domainItem), nil
}

func (r *queryResolver) Items(ctx context.Context) ([]model.Item, error) {
	rsctx := graphql.GetResolverContext(ctx)
	fmt.Println(rsctx)

	domainItems, err := r.item.GetItems(ctx)
	if err != nil {
		return nil, err
	}
	var items []model.Item
	for _, domainItem := range domainItems {
		items = append(items, *ToControllerItem(domainItem))
	}
	return items, nil
}

func (r *queryResolver) ItemHolder(ctx context.Context, id string) (*model.ItemHolder, error) {
	rsctx := graphql.GetResolverContext(ctx)
	fmt.Println(rsctx)

	res, err := r.itemHolder.GetItemHolder(ctx, id)
	if err != nil {
		return nil, err
	}
	return ToControllerItemHolder(res), nil
}

func (r *queryResolver) ItemHolders(ctx context.Context) ([]model.ItemHolder, error) {
	rsctx := graphql.GetResolverContext(ctx)
	fmt.Println(rsctx)

	results, err := r.itemHolder.GetItemHolders(ctx)
	if err != nil {
		return nil, err
	}
	var itemHolders []model.ItemHolder
	for _, res := range results {
		itemHolders = append(itemHolders, *ToControllerItemHolder(res))
	}
	return itemHolders, nil
}

/********************************************************************
 * Mutation
 */

func (r *mutationResolver) CreateItem(ctx context.Context, input ItemInput) (string, error) {
	rsctx := graphql.GetResolverContext(ctx)
	fmt.Println(rsctx)

	return r.item.CreateItem(ctx, ToCommandItemModel(input))
}

func (r *mutationResolver) CreateItemHolder(ctx context.Context, input ItemHolderInput) (string, error) {
	rsctx := graphql.GetResolverContext(ctx)
	fmt.Println(rsctx)

	return r.itemHolder.CreateItemHolder(ctx, ToCommandItemHolderModel(input))
}
