package database

import (
	"database/sql"
	"errors"
	"praticing/util"

	_ "github.com/lib/pq"
)

const (
	connectionString = "host=localhost port=5432 user=postgres dbname=postgres password=1234 sslmode=disable"
)

func ConnectDatabase() (*sql.DB, error) {
	connection, err := sql.Open("postgres", connectionString)

	util.CheckError(err, nil)

	return connection, checkConnection(connection)
}

func checkConnection(connection *sql.DB) error {
	if connection.Ping() != nil {
		return errors.New("570 - Can't connect to database")
	}
	return nil
}
