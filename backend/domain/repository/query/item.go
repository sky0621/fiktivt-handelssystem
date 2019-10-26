package query

type Item interface {
	GetItem(id string) *QueryItemModel
	GetItems() []*QueryItemModel
}

type QueryItemModel struct {
	ID         string
	Name       string
	Price      int64
	ItemHolder QueryItemHolderModel
}
