package main

import (
	"github.com/sky0621/fiktivt-handelssystem/config"
	"github.com/sky0621/fiktivt-handelssystem/domain/repository"
	"github.com/sky0621/fiktivt-handelssystem/driver"
)

func NewApp(cfg config.Config, rdb repository.Persistence, web driver.Web) App {
	return &AppImpl{
		cfg: cfg,
		rdb: rdb,
		web: web,
	}
}

type App interface {
	Start() error
	Shutdown()
}

type AppImpl struct {
	cfg config.Config
	rdb repository.Persistence
	web driver.Web
}

func (a *AppImpl) Start() error {
	// TODO:

	return a.web.Start()
}

func (a *AppImpl) Shutdown() {
	a.rdb.Close()
}
