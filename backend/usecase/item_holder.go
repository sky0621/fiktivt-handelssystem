package usecase

import (
	"github.com/sky0621/fiktivt-handelssystem/domain/repository/command"
	"github.com/sky0621/fiktivt-handelssystem/domain/repository/query"
)

func NewItemHolder() ItemHolder {
	return &itemHolder{}
}

type ItemHolder interface {
}

type itemHolder struct {
	itemHolderQuery   query.ItemHolder
	itemHolderCommand command.Item
}
