package controller

import (
	"github.com/google/uuid"
	"github.com/sky0621/fiktivt-handelssystem/adapter/controller/model"
	"github.com/sky0621/fiktivt-handelssystem/domain"
)

// MEMO: テスト時の置き換え用
var UniqueID = createUniqueID

// TODO: 置くべき場所は要検討
func createUniqueID() string {
	return uuid.New().String()
}

/*
 * From domain (for query)
 */

func ToControllerItem(m *domain.QueryItemModel) *model.Item {
	return &model.Item{
		ID:    m.ID,
		Name:  m.Name,
		Price: m.Price,
	}
}

func ToControllerItemHolder(m *domain.QueryItemHolderModel) *model.ItemHolder {
	var holdItems []model.Item
	for _, holdItem := range m.HoldItems {
		holdItems = append(holdItems, model.Item{
			ID:    holdItem.ID,
			Name:  holdItem.Name,
			Price: holdItem.Price,
		})
	}
	return &model.ItemHolder{
		ID:        m.ID,
		FirstName: m.FirstName,
		LastName:  m.LastName,
		Nickname:  m.Nickname,
	}
}

/*
 * From controller (for mutation)
 */

func ToCommandItemModel(input ItemInput) domain.CommandItemModel {
	return domain.CommandItemModel{
		ID:           UniqueID(),
		Name:         input.Name,
		Price:        input.Price,
		ItemHolderID: input.ItemHolderID,
	}
}

func ToCommandItemHolderModel(input ItemHolderInput) domain.CommandItemHolderModel {
	return domain.CommandItemHolderModel{
		ID:        UniqueID(),
		FirstName: input.FirstName,
		LastName:  input.LastName,
		Nickname:  input.Nickname,
	}
}
