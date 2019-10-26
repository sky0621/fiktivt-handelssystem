package command

import (
	"github.com/sky0621/fiktivt-handelssystem/domain/repository/command"
	"github.com/sky0621/fiktivt-handelssystem/driver"
)

func NewItemHolder(rdb driver.RDB) command.ItemHolder {
	return &itemHolder{rdb: rdb}
}

type itemHolder struct {
	rdb driver.RDB
}

func (i *itemHolder) CreateItemHolder(input command.CommandItemHolderModel) string {
	// FIXME:
	return "53f7978e-f16c-4002-9a8c-0df77d1145f0"
}
