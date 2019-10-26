package driver

import (
	"fmt"

	"github.com/sky0621/fiktivt-handelssystem/config"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type RDB interface {
	Close() error
}

func NewRDB(cfg config.Config) RDB {
	dsFormat := "dbname=%s user=%s password=%s sslmode=disable"
	dsn := fmt.Sprintf(dsFormat, cfg.RDBConfig.DBName, cfg.RDBConfig.User, cfg.RDBConfig.Password)
	dbWrapper, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		panic(err) // システム起動時なので
	}
	return &rdb{cfg: cfg, dbWrapper: dbWrapper}
}

type rdb struct {
	cfg       config.Config
	dbWrapper *sqlx.DB
}

func (r *rdb) Close() error {
	if r == nil {
		return nil
	}
	if r.dbWrapper == nil {
		return nil
	}
	return r.dbWrapper.Close()
}
