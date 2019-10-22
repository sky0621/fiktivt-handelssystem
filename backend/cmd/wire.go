//+build wireinject

package main

import (
	"github.com/sky0621/fiktivt-handelssystem/adapter/controller"
	"github.com/sky0621/fiktivt-handelssystem/driver"

	"github.com/sky0621/fiktivt-handelssystem/config"

	"github.com/google/wire"
)

var superSet = wire.NewSet(
	driver.NewRDB,
	driver.NewWeb,
	controller.NewRouter,
	NewApp,
)

func Initialize(cfg config.Config) App {
	wire.Build(superSet)
	return &AppImpl{}
}
