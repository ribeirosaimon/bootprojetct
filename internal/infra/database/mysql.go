package database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"sync"
)

var (
	_onceMysql       sync.Once
	_mysqlConnection *mysqlConnection
	err              error
)

type MysqlOption func(*mysqlConfig)

type mysqlConfig struct {
	url      string
	database string
}

type mysqlConnection struct {
	Connection *sql.DB
	url        string
}

func WithDatabase(database string) MysqlOption {
	return func(cfg *mysqlConfig) {
		cfg.database = database
	}
}

func WithUrl(url string) MysqlOption {
	return func(cfg *mysqlConfig) {
		cfg.url = url
	}
}

func defaultMysqlConfig() *mysqlConfig {
	return &mysqlConfig{
		url:      "user:userpassword@tcp(localhost:3306)",
		database: "encurtador",
	}
}

func NewMysqlConnection(opts ...MysqlOption) *mysqlConnection {
	_onceMysql.Do(func() {
		_mysqlConnection = &mysqlConnection{}
		cfg := defaultMysqlConfig()
		for _, o := range opts {
			o(cfg)
		}

		_mysqlConnection.Connection, err = sql.Open("mysql", fmt.Sprintf("%s/%s", cfg.url, cfg.database))
		if err != nil {
			panic(err)
		}
	})
	return _mysqlConnection
}
