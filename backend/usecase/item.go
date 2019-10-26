package usecase

import (
	"github.com/sky0621/fiktivt-handelssystem/domain/repository/command"
	"github.com/sky0621/fiktivt-handelssystem/domain/repository/query"
)

func NewItem(itemQuery query.Item, itemCommand command.Item) Item {
	return &item{
		itemQuery:   itemQuery,
		itemCommand: itemCommand,
	}
}

type Item interface {
	GetItem(id string) *query.QueryItemModel
	GetItems() []*query.QueryItemModel
}

type item struct {
	itemQuery   query.Item
	itemCommand command.Item
}

func (i *item) GetItem(id string) *query.QueryItemModel {
	return i.itemQuery.GetItem(id)
}

func (i *item) GetItems() []*query.QueryItemModel {
	return i.itemQuery.GetItems()
}
