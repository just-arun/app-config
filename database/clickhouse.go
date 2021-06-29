package database

import (
	"database/sql"
	"fmt"

	"gopkg.in/gorp.v1"
)

func (d *database) ClickHouseConnection(host string, port int) (dbmap *gorp.DbMap, err error) {
	connStr := fmt.Sprintf("http://%v:%v/postgres", host, port)
	db, err := sql.Open("clickhouse", connStr)
	if err != nil {
		return
	}
	if err = db.Ping(); err != nil {
		return
	}
	dbmap = &gorp.DbMap{Db: db, Dialect: gorp.PostgresDialect{}}
	return dbmap, nil
}
