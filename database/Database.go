package database

import (
	"database/sql"
	"praticing/errors"
	"praticing/util"

	_ "github.com/lib/pq"
)

const (
	connectionString = "host=localhost port=5432 user=postgres dbname=postgres password=1234 sslmode=disable"
)

func ConnectDatabase() (*sql.DB, error) {
	connection, err := sql.Open("postgres", connectionString)

	util.CheckError(err, nil)

	return connection, errors.CheckDatabaseConnection(connection)
}
