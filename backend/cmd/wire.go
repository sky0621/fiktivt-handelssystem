//+build wireinject

package main

import (
	"github.com/sky0621/fiktivt-handelssystem/adapter/controller"
	"github.com/sky0621/fiktivt-handelssystem/adapter/gateway"
	"github.com/sky0621/fiktivt-handelssystem/config"
	"github.com/sky0621/fiktivt-handelssystem/driver"
	"github.com/sky0621/fiktivt-handelssystem/system"
	"github.com/sky0621/fiktivt-handelssystem/usecase"

	"github.com/google/wire"
)

var superSet = wire.NewSet(
	// ロガー
	system.NewAppLogger,

	// Config => RDBコネクションプール
	driver.NewRDB,

	// RDBコネクションプール => domain層インタフェースをadapter層で実装
	gateway.NewItem,
	gateway.NewItemHolder,

	// domain層インタフェース => usecase層
	usecase.NewItem,
	usecase.NewItemHolder,

	// usecase層 => GraphQLリゾルバー
	controller.NewResolverRoot,

	// WebFramework
	driver.NewWeb,

	// システムそのもの
	NewApp,
)

func di(cfg config.Config) App {
	wire.Build(superSet)
	return &AppImpl{}
}
