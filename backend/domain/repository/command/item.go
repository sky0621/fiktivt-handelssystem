package command

type Item interface {
	CreateItem(input CommandItemModel) string
}

type CommandItemModel struct {
	ID           string
	Name         string
	Price        int64
	ItemHolderID string
}
