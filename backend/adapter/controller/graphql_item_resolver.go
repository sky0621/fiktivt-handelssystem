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
	lgr := r.logger.NewLogger("Resolver.Item")
	lgr.Info().Msg("call")

	return &itemResolver{r}
}

func (r *itemResolver) ItemHolder(ctx context.Context, obj *model.Item) (*model.ItemHolder, error) {
	lgr := r.logger.NewLogger("itemResolver.ItemHolder")
	lgr.Info().Msg("call")

	domainItemHolder, err := r.itemHolder.GetItemHolder(ctx, obj.ItemHolderID)
	if err != nil {
		lgr.Err(err)
		return nil, err
	}

	ret := ToControllerItemHolder(domainItemHolder)
	lgr.Info().Interface("model.ItemHolder", ret)

	return ret, nil
}

/********************************************************************
 * ItemHolderResolver
 */

type itemHolderResolver struct{ *Resolver }

func (r *Resolver) ItemHolder() ItemHolderResolver {
	lgr := r.logger.NewLogger("Resolver.ItemHolder")
	lgr.Info().Msg("call")

	return &itemHolderResolver{r}
}

func (r *itemHolderResolver) HoldItems(ctx context.Context, obj *model.ItemHolder) ([]model.Item, error) {
	lgr := r.logger.NewLogger("itemHolderResolver.HoldItems")
	lgr.Info().Msg("call")

	domainItems, err := r.item.GetItemsByItemHolderID(ctx, obj.ID)
	if err != nil {
		lgr.Err(err)
		return nil, err
	}
	var items []model.Item
	for _, domainItem := range domainItems {
		items = append(items, *ToControllerItem(domainItem))
	}

	lgr.Info().Interface("[]model.Item", items)

	return items, nil
}

/********************************************************************
 * Query
 */

func (r *queryResolver) Item(ctx context.Context, id string) (*model.Item, error) {
	lgr := r.logger.NewLogger("queryResolver.Item")
	lgr.Info().Msg("call")

	domainItem, err := r.item.GetItem(ctx, id, graphql.CollectAllFields(ctx))
	if err != nil {
		lgr.Err(err)
		return nil, err
	}

	ret := ToControllerItem(domainItem)
	lgr.Info().Interface("model.Item", ret)

	return ret, nil
}

func (r *queryResolver) Items(ctx context.Context) ([]model.Item, error) {
	lgr := r.logger.NewLogger("queryResolver.Items")
	lgr.Info().Msg("call")

	domainItems, err := r.item.GetItems(ctx)
	if err != nil {
		lgr.Err(err)
		return nil, err
	}
	var items []model.Item
	for _, domainItem := range domainItems {
		items = append(items, *ToControllerItem(domainItem))
	}

	lgr.Info().Interface("[]model.Item", items)

	return items, nil
}

func (r *queryResolver) ItemHolder(ctx context.Context, id string) (*model.ItemHolder, error) {
	lgr := r.logger.NewLogger("queryResolver.ItemHolder")
	lgr.Info().Msg("call")

	res, err := r.itemHolder.GetItemHolder(ctx, id)
	if err != nil {
		lgr.Err(err)
		return nil, err
	}

	ret := ToControllerItemHolder(res)
	lgr.Info().Interface("model.ItemHolder", ret)

	return ret, nil
}

func (r *queryResolver) ItemHolders(ctx context.Context) ([]model.ItemHolder, error) {
	lgr := r.logger.NewLogger("queryResolver.ItemHolders")
	lgr.Info().Msg("call")

	results, err := r.itemHolder.GetItemHolders(ctx)
	if err != nil {
		lgr.Err(err)
		return nil, err
	}
	var itemHolders []model.ItemHolder
	for _, res := range results {
		itemHolders = append(itemHolders, *ToControllerItemHolder(res))
	}

	lgr.Info().Interface("[]model.ItemHolder", itemHolders)

	return itemHolders, nil
}

func (r *queryResolver) ItemHoldersByCondition(ctx context.Context, searchWord *model.SearchWordCondition, itemHolder *model.SearchItemHolderCondition, first *int, after *string, sortCondition *model.SortCondition) (*model.ItemHolderConnection, error) {
	lgr := r.logger.NewLogger("queryResolver.ItemHoldersByCondition")
	lgr.Info().Msg("call")

	// domain層向けに変換
	searchWordModel := ToSearchWordConditionModel(searchWord)
	itemHolderModel := ToSearchItemHolderConditionModel(itemHolder)
	sortConditionModel := ToSortConditionModel(sortCondition)

	limit := 10 // デフォルト値は本来Config持ちかな
	if first != nil {
		limit = *first
	}

	result, allCount, err := r.itemHolder.GetItemHoldersByCondition(ctx, searchWordModel, itemHolderModel, limit, after, sortConditionModel)
	if err != nil {
		lgr.Err(err)
		return nil, err
	}

	// FIXME:
	fmt.Println(result)
	fmt.Println(allCount)

	return &model.ItemHolderConnection{
		TotalCount: allCount,
		Edges: []model.ItemHolderEdge{
			{Cursor: "", Node: &model.ItemHolder{
				ID:        "id0001",
				FirstName: "fn1",
				LastName:  "ln1",
				Nickname:  nil,
			}},
			{Cursor: "", Node: &model.ItemHolder{
				ID:        "id0002",
				FirstName: "fn2",
				LastName:  "ln2",
				Nickname:  nil,
			}},
		},
		PageInfo: &model.PageInfo{
			StartCursor: "id0003",
			EndCursor:   "ie0013",
			HasNextPage: true,
			HasPrevPage: true,
		},
	}, nil
}

/********************************************************************
 * Mutation
 */

func (r *mutationResolver) CreateItem(ctx context.Context, input ItemInput) (string, error) {
	lgr := r.logger.NewLogger("mutationResolver.CreateItem")
	lgr.Info().Msg("call")

	res, err := r.item.CreateItem(ctx, ToCommandItemModel(input))
	if err != nil {
		lgr.Err(err)
	}
	return res, nil
}

func (r *mutationResolver) CreateItemHolder(ctx context.Context, input ItemHolderInput) (string, error) {
	lgr := r.logger.NewLogger("mutationResolver.CreateItemHolder")
	lgr.Info().Msg("call")

	res, err := r.itemHolder.CreateItemHolder(ctx, ToCommandItemHolderModel(input))
	if err != nil {
		lgr.Err(err)
	}
	return res, nil
}
