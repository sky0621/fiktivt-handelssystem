package domain

type Item interface {
	GetItem(id string) (*QueryItemModel, error)
	GetItems() ([]*QueryItemModel, error)
	CreateItem(input CommandItemModel) (string, error)
}

type QueryItemModel struct {
	ID         string
	Name       string
	Price      int
	ItemHolder QueryItemHolderModel
}

type CommandItemModel struct {
	ID           string
	Name         string
	Price        int
	ItemHolderID string
}
