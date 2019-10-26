package usecase

import "github.com/sky0621/fiktivt-handelssystem/domain"

func NewItem() Item {
	return &item{}
}

type Item interface {
	GetItem(id string) *domain.ItemModel
}

type item struct {
}

func (i *item) GetItem(id string) *domain.ItemModel {

}
