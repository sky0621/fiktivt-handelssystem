package controller

import (
	"context"

	"github.com/sky0621/fiktivt-handelssystem/domain"

	"github.com/99designs/gqlgen/graphql"
)

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

func (r *queryResolver) Item(ctx context.Context, id string) (*Item, error) {
	lgr := r.logger.NewLogger("queryResolver.Item")
	lgr.Info().Msg("call")

	domainItem, err := r.item.GetItem(ctx, id, graphql.CollectAllFields(ctx))
	if err != nil {
		lgr.Err(err)
		return nil, err
	}

	ret := ToControllerItem(domainItem)
	lgr.Info().Interface("Item", ret)

	return ret, nil
}

func (r *queryResolver) Items(ctx context.Context) ([]Item, error) {
	lgr := r.logger.NewLogger("queryResolver.Items")
	lgr.Info().Msg("call")

	domainItems, err := r.item.GetItems(ctx)
	if err != nil {
		lgr.Err(err)
		return nil, err
	}
	var items []Item
	for _, domainItem := range domainItems {
		items = append(items, *ToControllerItem(domainItem))
	}

	lgr.Info().Interface("[]Item", items)

	return items, nil
}

func (r *queryResolver) ItemHolder(ctx context.Context, id string) (*ItemHolder, error) {
	lgr := r.logger.NewLogger("queryResolver.ItemHolder")
	lgr.Info().Msg("call")

	res, err := r.itemHolder.GetItemHolder(ctx, id)
	if err != nil {
		lgr.Err(err)
		return nil, err
	}

	ret := ToControllerItemHolder(res)
	lgr.Info().Interface("ItemHolder", ret)

	return ret, nil
}

func (r *queryResolver) ItemHolders(ctx context.Context) ([]ItemHolder, error) {
	lgr := r.logger.NewLogger("queryResolver.ItemHolders")
	lgr.Info().Msg("call")

	results, err := r.itemHolder.GetItemHolders(ctx)
	if err != nil {
		lgr.Err(err)
		return nil, err
	}
	var itemHolders []ItemHolder
	for _, res := range results {
		itemHolders = append(itemHolders, *ToControllerItemHolder(res))
	}

	lgr.Info().Interface("[]ItemHolder", itemHolders)

	return itemHolders, nil
}

func (r *queryResolver) ItemHoldersByCondition(ctx context.Context, baseCondition BaseCondition, addCondition *SearchItemHolderCondition) (*ItemHolderConnection, error) {
	lgr := r.logger.NewLogger("queryResolver.ItemHoldersByCondition")
	lgr.Info().Msg("call")

	// domain層向けに変換
	searchWordModel := ToSearchWordConditionModel(baseCondition.SearchWordCondition)
	itemHolderModel := ToSearchItemHolderConditionModel(addCondition)
	sortConditionModel := ToSortConditionModel(baseCondition.SortCondition)
	searchDirectionType := ToSearchDirectionType(baseCondition.SearchDirection)

	limit := 10 // デフォルト値は本来Config持ちかな
	if baseCondition.Limit != nil {
		limit = *baseCondition.Limit
	}
	limit++ // １件余分に取得することで「次ページ」、「前ページ」の存在有無判定に用いる

	itemHolders, allCount, err := r.itemHolder.GetItemHoldersByCondition(ctx,
		searchWordModel, itemHolderModel, sortConditionModel, searchDirectionType,
		limit, baseCondition.StartCursor, baseCondition.EndCursor)
	if err != nil {
		lgr.Err(err)
		return nil, err
	}

	var startCursor *string
	var endCursor *string
	hasPrevPage := false
	hasNextPage := false
	var nodes []*ItemHolder
	for idx, itemHolder := range itemHolders {
		// 本来の表示件数＋１件を取得しようとしているため、＋１件分はクライアントへは返却不要
		// 同時に、＋１件分まで取得できたということは、今回返却分よりもさらに「前ページ」ないし「次ページ」表示分の
		// データがあることを意味する。
		if idx == len(itemHolders)-1 {
			switch searchDirectionType {
			case domain.Prev:
				hasPrevPage = true
			case domain.Next:
				hasNextPage = true
			}
			continue
		}

		converted := ToControllerItemHolder(itemHolder)

		sortKey := "id"
		if sortConditionModel != nil {
			sortKey = sortConditionModel.SortKey
		}
		cursor := converted.GetCursor(sortKey)

		nodes = append(nodes, converted)

		if idx == 0 {
			startCursor = cursor
		}
		if idx == len(itemHolders)-2 {
			endCursor = cursor
		}
	}

	return &ItemHolderConnection{
		TotalCount: allCount,
		Nodes:      nodes,
		PageInfo: &PageInfo{
			StartCursor: startCursor,
			EndCursor:   endCursor,
			HasPrevPage: hasPrevPage,
			HasNextPage: hasNextPage,
		},
	}, nil
}

type itemResolver struct{ *Resolver }

func (r *Resolver) Item() ItemResolver {
	lgr := r.logger.NewLogger("Resolver.Item")
	lgr.Info().Msg("call")

	return &itemResolver{r}
}

func (r *itemResolver) ItemHolder(ctx context.Context, obj *Item) (*ItemHolder, error) {
	lgr := r.logger.NewLogger("itemResolver.ItemHolder")
	lgr.Info().Msg("call")

	domainItemHolder, err := r.itemHolder.GetItemHolder(ctx, obj.ItemHolderID)
	if err != nil {
		lgr.Err(err)
		return nil, err
	}

	ret := ToControllerItemHolder(domainItemHolder)
	lgr.Info().Interface("ItemHolder", ret)

	return ret, nil
}

type itemHolderResolver struct{ *Resolver }

func (r *Resolver) ItemHolder() ItemHolderResolver {
	lgr := r.logger.NewLogger("Resolver.ItemHolder")
	lgr.Info().Msg("call")

	return &itemHolderResolver{r}
}

func (r *itemHolderResolver) HoldItems(ctx context.Context, obj *ItemHolder) ([]Item, error) {
	lgr := r.logger.NewLogger("itemHolderResolver.HoldItems")
	lgr.Info().Msg("call")

	domainItems, err := r.item.GetItemsByItemHolderID(ctx, obj.ID)
	if err != nil {
		lgr.Err(err)
		return nil, err
	}
	var items []Item
	for _, domainItem := range domainItems {
		items = append(items, *ToControllerItem(domainItem))
	}

	lgr.Info().Interface("[]Item", items)

	return items, nil
}
