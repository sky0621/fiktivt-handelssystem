///+build wireinject

package main

import (
	"github.com/sky0621/fiktivt-handelssystem/adapter/controller"
	"github.com/sky0621/fiktivt-handelssystem/adapter/gateway/command"
	"github.com/sky0621/fiktivt-handelssystem/adapter/gateway/query"
	"github.com/sky0621/fiktivt-handelssystem/config"
	"github.com/sky0621/fiktivt-handelssystem/driver"
	"github.com/sky0621/fiktivt-handelssystem/usecase"

	"github.com/google/wire"
)

var superSet = wire.NewSet(
	// RDBコネクションプール
	driver.NewRDB,

	// DataAccess層
	command.NewItem,
	command.NewItemHolder,
	query.NewItem,
	query.NewItemHolder,

	// Usecase層
	usecase.NewItem,

	// InputportAdapter
	controller.NewGraphQLAdapter,

	// WebFramework
	driver.NewWeb,

	// システムそのもの
	NewApp,
)

func di(cfg config.Config) App {
	wire.Build(superSet)
	return &AppImpl{}
}
