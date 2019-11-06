package gateway

import (
	"context"
	"database/sql"
	"errors"
	"strings"

	"github.com/sky0621/fiktivt-handelssystem/system"

	"github.com/sky0621/fiktivt-handelssystem/adapter/gateway/model"

	"github.com/sky0621/fiktivt-handelssystem/domain"
	"github.com/sky0621/fiktivt-handelssystem/driver"
)

func NewItem(rdb driver.RDB, logger system.AppLogger) domain.Item {
	return &item{rdb: rdb, logger: logger}
}

type item struct {
	rdb    driver.RDB
	logger system.AppLogger
}

/********************************************************************
 * Query
 */

func (i *item) GetItem(ctx context.Context, id string, selectFields []string) (*domain.QueryItemModel, error) {
	i.logger.Log("call")

	sbQuery := strings.Builder{}
	sbQuery.WriteString("SELECT ")
	for _, selectField := range selectFields {
		if sbQuery.Len() > 7 {
			sbQuery.WriteString(", ")
		}
		switch selectField {
		case "id":
			sbQuery.WriteString("id")
		case "name":
			sbQuery.WriteString("name")
		case "price":
			sbQuery.WriteString("price")
		case "itemHolder":
			sbQuery.WriteString("item_holder_id")
		}
	}
	sbQuery.WriteString(" FROM item WHERE id = :id")
	q := sbQuery.String()
	i.logger.Log("query: " + q)

	stmt, err := i.rdb.GetDBWrapper().PrepareNamedContext(ctx, q)
	if err != nil {
		i.logger.Log(err.Error())
		return nil, err
	}

	res := &model.DBItem{}
	err = stmt.QueryRowxContext(ctx, map[string]interface{}{"id": id}).StructScan(res)
	if err != nil {
		i.logger.Log(err.Error())
		return nil, err
	}
	i.logger.Log(res.String())

	return &domain.QueryItemModel{
		ID:           res.ID,
		Name:         res.Name,
		Price:        res.Price,
		ItemHolderID: res.ItemHolderID,
	}, nil
}

func (i *item) GetItems(ctx context.Context) ([]*domain.QueryItemModel, error) {
	i.logger.Log("call")

	q := `SELECT id, name, price, item_holder_id FROM item`
	stmt, err := i.rdb.GetDBWrapper().PrepareNamedContext(ctx, q)
	if err != nil {
		i.logger.Log(err.Error())
		return nil, err
	}

	rows, err := stmt.QueryxContext(ctx, map[string]interface{}{})
	if err != nil {
		i.logger.Log(err.Error())
		return nil, err
	}

	var dests []*domain.QueryItemModel
	for rows.Next() {
		res := &model.DBItem{}
		err := rows.StructScan(&res)
		if err != nil {
			i.logger.Log(err.Error())
			return nil, err
		}
		dest := &domain.QueryItemModel{
			ID:           res.ID,
			Name:         res.Name,
			Price:        res.Price,
			ItemHolderID: res.ItemHolderID,
		}
		i.logger.Log(dest.String())
		dests = append(dests, dest)
	}
	return dests, nil
}

func (i *item) GetItemsByItemHolderID(ctx context.Context, itemHolderID string) ([]*domain.QueryItemModel, error) {
	i.logger.Log("call")

	q := `
		SELECT i.id, i.name, i.price, i.item_holder_id 
		FROM item i INNER JOIN item_holder_relation ih ON ih.item_id = i.id 
		WHERE ih.item_holder_id = :itemHolderID
	`
	stmt, err := i.rdb.GetDBWrapper().PrepareNamedContext(ctx, q)
	if err != nil {
		i.logger.Log(err.Error())
		return nil, err
	}

	rows, err := stmt.QueryxContext(ctx, map[string]interface{}{"itemHolderID": itemHolderID})
	if err != nil {
		i.logger.Log(err.Error())
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
	i.logger.Log("call")

	dbWrapper := i.rdb.GetDBWrapper()

	txx, err := dbWrapper.BeginTxx(ctx, &sql.TxOptions{})
	if err != nil {
		i.logger.Log(err.Error())
		return input.ID, err
	}
	defer func() {
		if err := txx.Rollback(); err != nil {
			i.logger.Log(err.Error())
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
		i.logger.Log(err.Error())
		return input.ID, err
	}

	itemSqlRes, err := itemStmt.ExecContext(ctx, map[string]interface{}{
		"id":           input.ID,
		"name":         input.Name,
		"price":        input.Price,
		"itemHolderID": input.ItemHolderID,
	})
	if err != nil {
		i.logger.Log(err.Error())
		return input.ID, err
	}
	rows, err := itemSqlRes.RowsAffected()
	if err != nil {
		i.logger.Log(err.Error())
		return input.ID, err
	}
	if rows != 1 {
		i.logger.Log(err.Error())
		return input.ID, errors.New("item affected rows != 1")
	}

	/*
	 * item_holder_relationテーブル登録
	 */
	itemHolderRelStmt, err := txx.PrepareNamedContext(ctx, `
		INSERT INTO item_holder_relation (item_id, item_holder_id) VALUES(:itemID, :itemHolderID)
	`)
	if err != nil {
		i.logger.Log(err.Error())
		return input.ID, err
	}

	itemHolderRelSqlRes, err := itemHolderRelStmt.ExecContext(ctx, map[string]interface{}{
		"itemID":       input.ID,
		"itemHolderID": input.ItemHolderID,
	})
	if err != nil {
		i.logger.Log(err.Error())
		return input.ID, err
	}
	rows, err = itemHolderRelSqlRes.RowsAffected()
	if err != nil {
		i.logger.Log(err.Error())
		return input.ID, err
	}
	if rows != 1 {
		i.logger.Log(err.Error())
		return input.ID, errors.New("item_holder_relation affected rows != 1")
	}

	err = txx.Commit()
	if err != nil {
		i.logger.Log(err.Error())
		return input.ID, err
	}

	return input.ID, nil
}
