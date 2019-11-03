package gateway

import (
	"context"
	"errors"
	"log"

	"github.com/sky0621/fiktivt-handelssystem/domain"
	"github.com/sky0621/fiktivt-handelssystem/driver"
)

func NewItemHolder(rdb driver.RDB) domain.ItemHolder {
	return &itemHolder{rdb: rdb}
}

type itemHolder struct {
	rdb driver.RDB
}

/********************************************************************
 * Query
 */

func (i *itemHolder) GetItemHolder(ctx context.Context, id string) (*domain.QueryItemHolderModel, error) {
	q := `SELECT id, first_name, last_name, nickname FROM item_holder WHERE id = :id`
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

// FIXME:
func (i *itemHolder) GetItemHolders(ctx context.Context) ([]*domain.QueryItemHolderModel, error) {
	one, err := i.GetItemHolder(ctx, "d4b8e9a5-1946-4fdd-8487-685babf319f7")
	if err != nil {
		return nil, err
	}
	return []*domain.QueryItemHolderModel{one}, nil
}

/********************************************************************
 * Mutation
 */

func (i *itemHolder) CreateItemHolder(ctx context.Context, input domain.CommandItemHolderModel) (string, error) {
	dbWrapper := i.rdb.GetDBWrapper()
	stmt, err := dbWrapper.PrepareNamedContext(ctx, `
		INSERT INTO item_holder (id, first_name, last_name, nickname) VALUES(:id, :firstName, :lastName, :nickname)
	`)
	if err != nil {
		// FIXME: log
		return input.ID, err
	}

	res, err := stmt.ExecContext(ctx, map[string]interface{}{
		"id":        input.ID,
		"firstName": input.FirstName,
		"lastName":  input.LastName,
		"nickname":  input.Nickname,
	})
	if err != nil {
		// FIXME: log
		return input.ID, err
	}
	rows, err := res.RowsAffected()
	if err != nil {
		// FIXME: log
		return input.ID, err
	}
	if rows != 1 {
		// FIXME: log
		return input.ID, errors.New("affected rows != 1")
	}

	return input.ID, nil
}
