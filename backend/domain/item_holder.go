package domain

import "context"

type ItemHolder interface {
	GetItemHolder(ctx context.Context, id string) (*QueryItemHolderModel, error)
	GetItemHolders(ctx context.Context) ([]*QueryItemHolderModel, error)
	CreateItemHolder(ctx context.Context, input CommandItemHolderModel) (string, error)
}

type QueryItemHolderModel struct {
	ID        string
	FirstName string
	LastName  string
	Nickname  *string
	HoldItems []QueryItemModel
}

type CommandItemHolderModel struct {
	ID        string
	FirstName string
	LastName  string
	Nickname  *string
}
