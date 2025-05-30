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

type BootProjectConnection interface {
	GetConnection() *sql.DB
}
type MysqlOption func(*mysqlConfig)

type mysqlConfig struct {
	url      string
	database string
	entry    string
	host     string
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

func WithHost(host string) MysqlOption {
	return func(cfg *mysqlConfig) {
		cfg.host = host
	}
}

func defaultMysqlConfig() *mysqlConfig {
	return &mysqlConfig{
		url:      "myuser:mypass@tcp",
		database: "shortener",
		entry:    "static/migrations",
		host:     "localhost:3306",
	}
}

func NewMysqlConnection(opts ...MysqlOption) *mysqlConnection {
	_onceMysql.Do(func() {
		_mysqlConnection = &mysqlConnection{}
		cfg := defaultMysqlConfig()
		for _, o := range opts {
			o(cfg)
		}
		host := fmt.Sprintf("%s(%s)/%s", cfg.url, cfg.host, cfg.database)
		_mysqlConnection.Connection, err = sql.Open("mysql", host)
		if err != nil {
			panic(err)
		}
	})
	return _mysqlConnection
}

func (m *mysqlConnection) GetConnection() *sql.DB {
	return m.Connection
}
