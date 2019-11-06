package gateway

import (
	"context"
	"errors"
	"log"

	"github.com/sky0621/fiktivt-handelssystem/system"

	"github.com/sky0621/fiktivt-handelssystem/adapter/gateway/model"

	"github.com/sky0621/fiktivt-handelssystem/domain"
	"github.com/sky0621/fiktivt-handelssystem/driver"
)

func NewItemHolder(rdb driver.RDB, logger system.AppLogger) domain.ItemHolder {
	return &itemHolder{rdb: rdb, logger: logger}
}

type itemHolder struct {
	rdb    driver.RDB
	logger system.AppLogger
}

/********************************************************************
 * Query
 */

func (i *itemHolder) GetItemHolder(ctx context.Context, id string) (*domain.QueryItemHolderModel, error) {
	i.logger.Log("call")

	q := `SELECT id, first_name, last_name, nickname FROM item_holder WHERE id = :id`
	stmt, err := i.rdb.GetDBWrapper().PrepareNamedContext(ctx, q)
	if err != nil {
		i.logger.Log(err.Error())
		return nil, err
	}

	res := make(map[string]interface{})
	err = stmt.QueryRowxContext(ctx, map[string]interface{}{"id": id}).MapScan(res)
	if err != nil {
		i.logger.Log(err.Error())
		return nil, err
	}
	log.Println(res)

	// FIXME: とりあえずエラーハンドリングも型安全も考慮せず適当にマッピング
	resID := res["id"].(string)
	resFirstName := res["first_name"].(string)
	resLastName := res["last_name"].(string)
	resNickname := res["nickname"].(string)
	return &domain.QueryItemHolderModel{
		ID:        resID,
		FirstName: resFirstName,
		LastName:  resLastName,
		Nickname:  &resNickname,
	}, nil
}

func (i *itemHolder) GetItemHolders(ctx context.Context) ([]*domain.QueryItemHolderModel, error) {
	i.logger.Log("call")

	q := `SELECT id, first_name, last_name, nickname FROM item_holder`
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

	var dests []*domain.QueryItemHolderModel
	for rows.Next() {
		res := &model.DBItemHolder{}
		err := rows.StructScan(&res)
		if err != nil {
			i.logger.Log(err.Error())
			return nil, err
		}
		dests = append(dests, &domain.QueryItemHolderModel{
			ID:        res.ID,
			FirstName: res.FirstName,
			LastName:  res.LastName,
			Nickname:  &res.Nickname,
		})
	}
	return dests, nil
}

/********************************************************************
 * Mutation
 */

func (i *itemHolder) CreateItemHolder(ctx context.Context, input domain.CommandItemHolderModel) (string, error) {
	i.logger.Log("call")

	dbWrapper := i.rdb.GetDBWrapper()
	stmt, err := dbWrapper.PrepareNamedContext(ctx, `
		INSERT INTO item_holder (id, first_name, last_name, nickname) VALUES(:id, :firstName, :lastName, :nickname)
	`)
	if err != nil {
		i.logger.Log(err.Error())
		return input.ID, err
	}

	res, err := stmt.ExecContext(ctx, map[string]interface{}{
		"id":        input.ID,
		"firstName": input.FirstName,
		"lastName":  input.LastName,
		"nickname":  input.Nickname,
	})
	if err != nil {
		i.logger.Log(err.Error())
		return input.ID, err
	}
	rows, err := res.RowsAffected()
	if err != nil {
		i.logger.Log(err.Error())
		return input.ID, err
	}
	if rows != 1 {
		i.logger.Log(err.Error())
		return input.ID, errors.New("affected rows != 1")
	}

	return input.ID, nil
}
