package gateway

import (
	"github.com/sky0621/fiktivt-handelssystem/domain"
	"github.com/sky0621/fiktivt-handelssystem/driver"
)

func NewItem(rdb driver.RDB) domain.Item {
	return &item{rdb: rdb}
}

type item struct {
	rdb driver.RDB
}

func (i *item) GetItem(id string) (*domain.QueryItemModel, error) {
	// FIXME:
	nickname := "所有者１のニックネーム"
	return &domain.QueryItemModel{
		ID:    id,
		Name:  "商品１",
		Price: 1000,
		ItemHolder: domain.QueryItemHolderModel{
			ID:        "d4b8e9a5-1946-4fdd-8487-685babf319f7",
			Name:      "所有者１",
			Nickname:  &nickname,
			HoldItems: nil,
		},
	}, nil
}

func (i *item) GetItems() ([]*domain.QueryItemModel, error) {
	one, err := i.GetItem("97a835cd-f99a-4bf8-8928-13a5fe7d6552")
	if err != nil {
		return nil, err
	}
	return []*domain.QueryItemModel{one}, nil
}

func (i *item) CreateItem(input domain.CommandItemModel) (string, error) {
	// FIXME:
	return "4275a724-d693-42a5-93b6-3ffca0c3da61", nil
}
