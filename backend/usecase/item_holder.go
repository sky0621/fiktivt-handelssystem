package usecase

import (
	"github.com/sky0621/fiktivt-handelssystem/domain"
)

func NewItemHolder(itemHolderDomain domain.ItemHolder) ItemHolder {
	return &itemHolder{itemHolderDomain: itemHolderDomain}
}

type ItemHolder interface {
	GetItemHolder(id string) (*domain.QueryItemHolderModel, error)
	GetItemHolders() ([]*domain.QueryItemHolderModel, error)
	CreateItemHolder(input domain.CommandItemHolderModel) (string, error)
}

type itemHolder struct {
	itemHolderDomain domain.ItemHolder
}

func (i *itemHolder) GetItemHolder(id string) (*domain.QueryItemHolderModel, error) {
	return i.itemHolderDomain.GetItemHolder(id)
}

func (i *itemHolder) GetItemHolders() ([]*domain.QueryItemHolderModel, error) {
	return i.itemHolderDomain.GetItemHolders()
}

func (i *itemHolder) CreateItemHolder(input domain.CommandItemHolderModel) (string, error) {
	return i.itemHolderDomain.CreateItemHolder(input)
}
