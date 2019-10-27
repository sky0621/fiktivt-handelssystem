package controller

import (
	"github.com/sky0621/fiktivt-handelssystem/domain"
)

/*
 * From domain (for query)
 */

func ToControllerItem(m *domain.QueryItemModel) *Item {
	return &Item{
		ID:         m.ID,
		Name:       m.Name,
		Price:      m.Price,
		ItemHolder: ToControllerItemHolder(&m.ItemHolder),
	}
}

func ToControllerItemHolder(m *domain.QueryItemHolderModel) *ItemHolder {
	var holdItems []Item
	for _, holdItem := range m.HoldItems {
		holdItems = append(holdItems, Item{
			ID:    holdItem.ID,
			Name:  holdItem.Name,
			Price: holdItem.Price,
		})
	}
	return &ItemHolder{
		ID:        m.ID,
		Name:      m.Name,
		Nickname:  m.Nickname,
		HoldItems: holdItems,
	}
}

/*
 * From controller (for mutation)
 */

func ToCommandItemModel(input ItemInput) domain.CommandItemModel {
	return domain.CommandItemModel{
		Name:         input.Name,
		Price:        input.Price,
		ItemHolderID: input.ItemHolderID,
	}
}

func ToCommandItemHolderModel(input ItemHolderInput) domain.CommandItemHolderModel {
	return domain.CommandItemHolderModel{
		Name:     input.Name,
		Nickname: input.Nickname,
	}
}
