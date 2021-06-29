package database

import (
	"database/sql"

	"gopkg.in/gorp.v1"
)

func CockroachConnection(connectionUrl string) (*gorp.DbMap, error) {
	db, err := sql.Open("postgres", connectionUrl)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	if _, err := db.Exec("CREATE DATABASE IF NOT EXISTS bank;"); err != nil {
		return nil, err
	}
	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.PostgresDialect{}}
	return dbmap, nil
}
