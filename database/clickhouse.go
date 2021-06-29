package database

import (
	"database/sql"

	"gopkg.in/gorp.v1"
)

// connection string: http://host:port/postgres
func ClickHouseConnection(connectionString string) (dbmap *gorp.DbMap, err error) {
	db, err := sql.Open("clickhouse", connectionString)
	if err != nil {
		return
	}
	if err = db.Ping(); err != nil {
		return
	}
	dbmap = &gorp.DbMap{Db: db, Dialect: gorp.PostgresDialect{}}
	return dbmap, nil
}
