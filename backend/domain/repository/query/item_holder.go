package query

type ItemHolder interface {
	GetItemHolder(id string) *QueryItemHolderModel
}

type QueryItemHolderModel struct {
	ID        string
	Name      string
	Nickname  *string
	HoldItems []QueryItemModel
}
