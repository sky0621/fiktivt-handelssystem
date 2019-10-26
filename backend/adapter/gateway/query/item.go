package query

import (
	"github.com/sky0621/fiktivt-handelssystem/domain/repository/query"
	"github.com/sky0621/fiktivt-handelssystem/driver"
)

func NewItem(rdb driver.RDB) query.Item {
	return &item{rdb: rdb}
}

type item struct {
	rdb driver.RDB
}

func (i *item) GetItem(id string) *query.QueryItemModel {
	// FIXME:
	nickname := "所有者１のニックネーム"
	return &query.QueryItemModel{
		ID:    "97a835cd-f99a-4bf8-8928-13a5fe7d6552",
		Name:  "商品１",
		Price: 1000,
		ItemHolder: query.QueryItemHolderModel{
			ID:        "d4b8e9a5-1946-4fdd-8487-685babf319f7",
			Name:      "所有者１",
			Nickname:  &nickname,
			HoldItems: nil,
		},
	}
}
