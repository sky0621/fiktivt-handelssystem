package command

import (
	"github.com/sky0621/fiktivt-handelssystem/domain/repository/command"
	"github.com/sky0621/fiktivt-handelssystem/driver"
)

func NewItem(rdb driver.RDB) command.Item {
	return &item{rdb: rdb}
}

type item struct {
	rdb driver.RDB
}

func (i *item) CreateItem(input command.CommandItemModel) string {
	// FIXME:
	return "4275a724-d693-42a5-93b6-3ffca0c3da61"
}
