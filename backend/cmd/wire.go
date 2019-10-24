//+build wireinject

package main

import (
	"github.com/sky0621/fiktivt-handelssystem/adapter/controller"
	"github.com/sky0621/fiktivt-handelssystem/domain"
	"github.com/sky0621/fiktivt-handelssystem/domain/repository"
	"github.com/sky0621/fiktivt-handelssystem/driver"
	"github.com/sky0621/fiktivt-handelssystem/usecase"

	"github.com/sky0621/fiktivt-handelssystem/config"

	"github.com/google/wire"
)

var superSet = wire.NewSet(
	// RDBコネクションプール
	driver.NewRDB,

	// DataAccess層
	repository.NewOrganization,
	repository.NewUser,
	repository.NewOrder,
	repository.NewOrderDetail,
	repository.NewInstruction,

	// BusinessLogic層
	domain.NewOrganization,
	domain.NewOrder,

	// Usecase層
	usecase.NewOrder,

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
