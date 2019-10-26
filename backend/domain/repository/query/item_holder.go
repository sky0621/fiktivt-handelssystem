package repository

func NewItemHolderQuery() ItemHolderQuery {
	return &itemHolderQuery{}
}

type ItemHolderQuery interface {
	GetItemHolder(id string) *ItemHolderModel
}

type itemHolderQuery struct {
}

func (q *itemHolderQuery) GetItemHolder(id string) *ItemHolderModel {

}

type ItemHolderModel struct {
}
