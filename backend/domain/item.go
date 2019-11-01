package domain

import "context"

type Item interface {
	GetItem(ctx context.Context, id string, selectFields []string) (*QueryItemModel, error)
	GetItems(ctx context.Context) ([]*QueryItemModel, error)
	GetItemsByItemHolderID(ctx context.Context, itemHolderID string) ([]*QueryItemModel, error)
	CreateItem(ctx context.Context, input CommandItemModel) (string, error)
}

type QueryItemModel struct {
	ID    string
	Name  string
	Price int
}

type CommandItemModel struct {
	ID           string
	Name         string
	Price        int
	ItemHolderID string
}
