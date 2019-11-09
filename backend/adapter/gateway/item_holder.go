package gateway

import (
	"context"
	"errors"

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
	lgr := i.logger.NewLogger("itemHolder.GetItemHolder")
	lgr.Info().Msg("call")

	q := `SELECT id, first_name, last_name, nickname FROM item_holder WHERE id = :id`
	stmt, err := i.rdb.GetDBWrapper().PrepareNamedContext(ctx, q)
	if err != nil {
		lgr.Err(err)
		return nil, err
	}

	res := &model.DBItemHolder{}
	err = stmt.QueryRowxContext(ctx, map[string]interface{}{"id": id}).StructScan(res)
	if err != nil {
		lgr.Err(err)
		return nil, err
	}
	lgr.Info().Str("model.DBItemHolder", res.String()).Send()

	return &domain.QueryItemHolderModel{
		ID:        res.ID,
		FirstName: res.FirstName,
		LastName:  res.LastName,
		Nickname:  &res.Nickname,
	}, nil
}

func (i *itemHolder) GetItemHolders(ctx context.Context) ([]*domain.QueryItemHolderModel, error) {
	lgr := i.logger.NewLogger("itemHolder.GetItemHolders")
	lgr.Info().Msg("call")

	q := `SELECT id, first_name, last_name, nickname FROM item_holder`
	stmt, err := i.rdb.GetDBWrapper().PrepareNamedContext(ctx, q)
	if err != nil {
		lgr.Err(err)
		return nil, err
	}

	rows, err := stmt.QueryxContext(ctx, map[string]interface{}{})
	if err != nil {
		lgr.Err(err)
		return nil, err
	}

	var dests []*domain.QueryItemHolderModel
	for rows.Next() {
		res := &model.DBItemHolder{}
		err := rows.StructScan(&res)
		if err != nil {
			lgr.Err(err)
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

func (i *itemHolder) GetItemHoldersByCondition(ctx context.Context,
	searchWordCondition *domain.SearchWordConditionModel,
	itemHolderCondition *domain.SearchItemHolderConditionModel,
	sortCondition *domain.SortConditionModel,
	searchDirectionType domain.SearchDirection,
	limit int, startCursor *string, endCursor *string,
) ([]*domain.QueryItemHolderModel, int, error) {
	// FIXME:
	return []*domain.QueryItemHolderModel{}, 0, nil
}

/********************************************************************
 * Mutation
 */

func (i *itemHolder) CreateItemHolder(ctx context.Context, input domain.CommandItemHolderModel) (string, error) {
	lgr := i.logger.NewLogger("itemHolder.CreateItemHolder")
	lgr.Info().Msg("call")

	dbWrapper := i.rdb.GetDBWrapper()
	stmt, err := dbWrapper.PrepareNamedContext(ctx, `
		INSERT INTO item_holder (id, first_name, last_name, nickname) VALUES(:id, :firstName, :lastName, :nickname)
	`)
	if err != nil {
		lgr.Err(err)
		return input.ID, err
	}

	res, err := stmt.ExecContext(ctx, map[string]interface{}{
		"id":        input.ID,
		"firstName": input.FirstName,
		"lastName":  input.LastName,
		"nickname":  input.Nickname,
	})
	if err != nil {
		lgr.Err(err)
		return input.ID, err
	}
	rows, err := res.RowsAffected()
	if err != nil {
		lgr.Err(err)
		return input.ID, err
	}
	if rows != 1 {
		lgr.Err(err)
		return input.ID, errors.New("affected rows != 1")
	}

	return input.ID, nil
}
