package domain

import "context"

type Item interface {
	GetItem(ctx context.Context, id string) (*QueryItemModel, error)
	GetItems(ctx context.Context) ([]*QueryItemModel, error)
	CreateItem(ctx context.Context, input CommandItemModel) (string, error)
}

type QueryItemModel struct {
	ID    string
	Name  string
	Price int
	//ItemHolder QueryItemHolderModel
}

type CommandItemModel struct {
	ID           string
	Name         string
	Price        int
	ItemHolderID string
}
