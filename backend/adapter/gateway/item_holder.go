package gateway

import (
	"context"

	"github.com/sky0621/fiktivt-handelssystem/domain"
	"github.com/sky0621/fiktivt-handelssystem/driver"
)

func NewItemHolder(rdb driver.RDB) domain.ItemHolder {
	return &itemHolder{rdb: rdb}
}

type itemHolder struct {
	rdb driver.RDB
}

func (i *itemHolder) GetItemHolder(ctx context.Context, id string) (*domain.QueryItemHolderModel, error) {
	// FIXME:
	nickname := "所有者１のニックネーム"
	return &domain.QueryItemHolderModel{
		ID:       "d4b8e9a5-1946-4fdd-8487-685babf319f7",
		Name:     "所有者１",
		Nickname: &nickname,
		HoldItems: []domain.QueryItemModel{
			{
				ID:    "97a835cd-f99a-4bf8-8928-13a5fe7d6552",
				Name:  "商品１",
				Price: 1000,
			},
		},
	}, nil
}

func (i *itemHolder) GetItemHolders(ctx context.Context) ([]*domain.QueryItemHolderModel, error) {
	one, err := i.GetItemHolder(ctx, "d4b8e9a5-1946-4fdd-8487-685babf319f7")
	if err != nil {
		return nil, err
	}
	return []*domain.QueryItemHolderModel{one}, nil
}

func (i *itemHolder) CreateItemHolder(ctx context.Context, input domain.CommandItemHolderModel) (string, error) {
	dbWrapper := i.rdb.GetDBWrapper()
	dbWrapper.PrepareNamedContext(ctx, `
		INSERT INTO item_holder
	`)

	// FIXME:
	return "53f7978e-f16c-4002-9a8c-0df77d1145f0", nil
}
