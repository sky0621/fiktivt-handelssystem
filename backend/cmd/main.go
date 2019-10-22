package main

import (
	"os"

	"github.com/sky0621/fiktivt-handelssystem/config"
)

func main() {
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

	app := Initialize(cfg)
	defer app.Shutdown()
	app.Start()
}
