package controller

import (
	"github.com/sky0621/fiktivt-handelssystem/usecase"
)

func NewGraphQLAdapter(itemUsecase usecase.Item) GraphQLAdapter {
	return &graphQLAdapter{
		itemUsecase: itemUsecase,
	}
}

type GraphQLAdapter interface {
	GetItemUsecase() usecase.Item
}

type graphQLAdapter struct {
	itemUsecase usecase.Item
	// TODO: ユースケース追加！
}

func (a *graphQLAdapter) GetItemUsecase() usecase.Item {
	return a.itemUsecase
}
