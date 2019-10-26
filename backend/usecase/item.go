package usecase

import (
	"github.com/sky0621/fiktivt-handelssystem/domain"
)

func NewItem(itemDomain domain.Item) Item {
	return &item{
		itemDomain: itemDomain,
	}
}

type Item interface {
	GetItem(id string) (*domain.QueryItemModel, error)
	GetItems() ([]*domain.QueryItemModel, error)
	CreateItem(input domain.CommandItemModel) (string, error)
}

type item struct {
	itemDomain domain.Item
}

func (i *item) GetItem(id string) (*domain.QueryItemModel, error) {
	return i.itemDomain.GetItem(id)
}

func (i *item) GetItems() ([]*domain.QueryItemModel, error) {
	return i.itemDomain.GetItems()
}

func (i *item) CreateItem(input domain.CommandItemModel) (string, error) {
	return i.itemDomain.CreateItem(input)
}
