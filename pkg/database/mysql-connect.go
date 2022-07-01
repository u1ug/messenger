package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

func NewDefaultConnection(dataSouceName string) (*sql.DB, error) {
	connectionPool, err := sql.Open("mysql", dataSouceName)
	if err != nil {
		panic(err)
	}
	connectionPool.SetConnMaxLifetime(time.Minute * 10)
	connectionPool.SetMaxIdleConns(50)
	connectionPool.SetMaxOpenConns(25)
	connectionPool.SetConnMaxLifetime(time.Minute * 5)
	return connectionPool, err
}
