package gateway

import (
	"context"
	"database/sql"
	"errors"
	"log"

	"github.com/sky0621/fiktivt-handelssystem/domain"
	"github.com/sky0621/fiktivt-handelssystem/driver"
)

func NewItem(rdb driver.RDB) domain.Item {
	return &item{rdb: rdb}
}

type item struct {
	rdb driver.RDB
}

func (i *item) GetItem(ctx context.Context, id string) (*domain.QueryItemModel, error) {
	q := `
		SELECT id, name, price, item_holder_id FROM item WHERE id = :id
	`
	stmt, err := i.rdb.GetDBWrapper().PrepareNamedContext(ctx, q)
	if err != nil {
		return nil, err
	}

	res := make(map[string]interface{})
	err = stmt.QueryRowxContext(ctx, map[string]interface{}{"id": id}).MapScan(res)
	if err != nil {
		return nil, err
	}
	log.Println(res)

	// FIXME: とりあえずエラーハンドリングも型安全も考慮せず適当にマッピング
	resID := res["id"].(string)
	resName := res["name"].(string)
	resPrice, ok := res["price"].(int64)
	if !ok {
		return nil, err
	}
	return &domain.QueryItemModel{
		ID:    resID,
		Name:  resName,
		Price: int(resPrice),
		//ItemHolder: domain.QueryItemHolderModel{
		//	ID:        "d4b8e9a5-1946-4fdd-8487-685babf319f7",
		//	Name:      "所有者１",
		//	Nickname:  &nickname,
		//	HoldItems: nil,
		//},
	}, nil
}

func (i *item) GetItems(ctx context.Context) ([]*domain.QueryItemModel, error) {
	one, err := i.GetItem(ctx, "97a835cd-f99a-4bf8-8928-13a5fe7d6552")
	if err != nil {
		return nil, err
	}
	return []*domain.QueryItemModel{one}, nil
}

func (i *item) CreateItem(ctx context.Context, input domain.CommandItemModel) (string, error) {
	dbWrapper := i.rdb.GetDBWrapper()

	txx, err := dbWrapper.BeginTxx(ctx, &sql.TxOptions{})
	if err != nil {
		// FIXME: log
		return input.ID, err
	}
	defer func() {
		if err := txx.Rollback(); err != nil {
			// FIXME: log
			log.Println(err)
			return
		}
	}()

	/*
	 * itemテーブル登録
	 */
	itemStmt, err := txx.PrepareNamedContext(ctx, `
		INSERT INTO item (id, name, price, item_holder_id) VALUES(:id, :name, :price, :itemHolderID)
	`)
	if err != nil {
		// FIXME: log
		return input.ID, err
	}

	itemSqlRes, err := itemStmt.ExecContext(ctx, map[string]interface{}{
		"id":           input.ID,
		"name":         input.Name,
		"price":        input.Price,
		"itemHolderID": input.ItemHolderID,
	})
	if err != nil {
		// FIXME: log
		return input.ID, err
	}
	rows, err := itemSqlRes.RowsAffected()
	if err != nil {
		// FIXME: log
		return input.ID, err
	}
	if rows != 1 {
		// FIXME: log
		return input.ID, errors.New("item affected rows != 1")
	}

	/*
	 * item_holder_relationテーブル登録
	 */
	itemHolderRelStmt, err := txx.PrepareNamedContext(ctx, `
		INSERT INTO item_holder_relation (item_id, item_holder_id) VALUES(:itemIDx, :itemHolderID)
	`)
	if err != nil {
		// FIXME: log
		return input.ID, err
	}

	itemHolderRelSqlRes, err := itemHolderRelStmt.ExecContext(ctx, map[string]interface{}{
		"itemID":       input.ID,
		"itemHolderID": input.ItemHolderID,
	})
	if err != nil {
		// FIXME: log
		return input.ID, err
	}
	rows, err = itemHolderRelSqlRes.RowsAffected()
	if err != nil {
		// FIXME: log
		return input.ID, err
	}
	if rows != 1 {
		// FIXME: log
		return input.ID, errors.New("item_holder_relation affected rows != 1")
	}

	err = txx.Commit()
	if err != nil {
		// FIXME: log
		return input.ID, err
	}

	return input.ID, nil
}
