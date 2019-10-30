package gateway

import (
	"context"
	"database/sql"
	"errors"
	"log"

	"github.com/sky0621/fiktivt-handelssystem/adapter/gateway/model"

	"github.com/sky0621/fiktivt-handelssystem/domain"
	"github.com/sky0621/fiktivt-handelssystem/driver"
)

func NewItem(rdb driver.RDB) domain.Item {
	return &item{rdb: rdb}
}

type item struct {
	rdb driver.RDB
}

/********************************************************************
 * Query
 */

func (i *item) GetItem(ctx context.Context, id string) (*domain.QueryItemModel, error) {
	q := `SELECT id, name, price, item_holder_id FROM item WHERE id = :id`
	stmt, err := i.rdb.GetDBWrapper().PrepareNamedContext(ctx, q)
	if err != nil {
		return nil, err
	}

	res := &model.DBItem{}
	err = stmt.QueryRowxContext(ctx, map[string]interface{}{"id": id}).StructScan(res)
	if err != nil {
		return nil, err
	}
	log.Println(res)

	return &domain.QueryItemModel{
		ID:    res.ID,
		Name:  res.Name,
		Price: res.Price,
	}, nil
}

// FIXME:
func (i *item) GetItems(ctx context.Context) ([]*domain.QueryItemModel, error) {
	q := `SELECT id, name, price, item_holder_id FROM item`
	stmt, err := i.rdb.GetDBWrapper().PrepareNamedContext(ctx, q)
	if err != nil {
		return nil, err
	}

	rows, err := stmt.QueryxContext(ctx, map[string]interface{}{})
	if err != nil {
		return nil, err
	}

	var dests []*domain.QueryItemModel
	for rows.Next() {
		res := &model.DBItem{}
		err := rows.StructScan(&res)
		if err != nil {
			return nil, err
		}
		dests = append(dests, &domain.QueryItemModel{
			ID:    res.ID,
			Name:  res.Name,
			Price: res.Price,
		})
	}
	return dests, nil
}

func (i *item) GetItemsByItemHolderID(ctx context.Context, itemHolderID string) ([]*domain.QueryItemModel, error) {
	q := `
		SELECT i.id, i.name, i.price, i.item_holder_id 
		FROM item i INNER JOIN item_holder_relation ih ON ih.item_id = i.id 
		WHERE ih.item_holder_id = :itemHolderID
	`
	stmt, err := i.rdb.GetDBWrapper().PrepareNamedContext(ctx, q)
	if err != nil {
		return nil, err
	}

	rows, err := stmt.QueryxContext(ctx, map[string]interface{}{"itemHolderID": itemHolderID})
	if err != nil {
		return nil, err
	}

	var dests []*domain.QueryItemModel
	for rows.Next() {
		res := &model.DBItem{}
		err := rows.StructScan(&res)
		if err != nil {
			return nil, err
		}
		dests = append(dests, &domain.QueryItemModel{
			ID:    res.ID,
			Name:  res.Name,
			Price: res.Price,
		})
	}
	return dests, nil
}

/********************************************************************
 * Mutation
 */

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
		INSERT INTO item_holder_relation (item_id, item_holder_id) VALUES(:itemID, :itemHolderID)
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
