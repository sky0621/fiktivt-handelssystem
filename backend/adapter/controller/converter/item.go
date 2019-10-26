package converter

import (
	"github.com/sky0621/fiktivt-handelssystem/adapter/controller"
	"github.com/sky0621/fiktivt-handelssystem/domain"
)

func ToItemHolder(m *domain.QueryItemHolderModel) *controller.ItemHolder {
	return &controller.ItemHolder{
		ID:        m.ID,
		Name:      m.Name,
		Nickname:  m.Nickname,
		HoldItems: m.HoldItems,
	}
}

func ToCommandItemModel(input controller.ItemInput) domain.CommandItemModel {
	return domain.CommandItemModel{
		Name:         input.Name,
		Price:        input.Price,
		ItemHolderID: input.ItemHolderID,
	}
}

func ToCommandItemHolderModel(input controller.ItemHolderInput) domain.CommandItemHolderModel {
	return domain.CommandItemHolderModel{
		Name:     input.Name,
		Nickname: input.Nickname,
	}
}
