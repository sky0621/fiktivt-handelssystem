package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/sky0621/fiktivt-handelssystem/config"
)

// TODO: 既存である？
type exitCode int

const (
	normalEnd   = 0
	abnormalEnd = -1
)

func main() {
	os.Exit(int(execMain()))
}

func execMain() exitCode {
	cfg := config.Config{
		RDBConfig: config.RDBConfig{
			DBName:   os.Getenv("FIKTIVT_RDB_DBNAME"),
			User:     os.Getenv("FIKTIVT_RDB_USER"),
			Password: os.Getenv("FIKTIVT_RDB_PASSWORD"),
		},
		WebConfig: config.WebConfig{ListenPort: os.Getenv("FIKTIVT_WEB_LISTENPORT")},
		AppConfig: config.AppConfig{
			// FIXME
		},
	}

	app := di(cfg)
	defer app.Shutdown()

	// OSシグナル受信したらグレースフルシャットダウン
	go func() {
		q := make(chan os.Signal)
		signal.Notify(q, os.Interrupt, os.Kill, syscall.SIGTERM)
		<-q
		app.Shutdown()
		os.Exit(int(abnormalEnd))
	}()

	if err := app.Start(); err != nil {
		fmt.Errorf("the application failed to start")
		return abnormalEnd
	}
	return normalEnd
}
