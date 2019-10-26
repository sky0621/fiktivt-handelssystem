package domain

type ItemHolder interface {
	GetItemHolder(id string) (*QueryItemHolderModel, error)
	GetItemHolders() ([]*QueryItemHolderModel, error)
	CreateItemHolder(input CommandItemHolderModel) (string, error)
}

type QueryItemHolderModel struct {
	ID        string
	Name      string
	Nickname  *string
	HoldItems []QueryItemModel
}

type CommandItemHolderModel struct {
	ID       string
	Name     string
	Nickname *string
}
