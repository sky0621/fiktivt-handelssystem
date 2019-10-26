package repository

func NewItemQuery() ItemQuery {
	return &itemQuery{}
}

type ItemQuery interface {
	GetItem(id string)
}

type itemQuery struct {
}
