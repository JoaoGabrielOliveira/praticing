package config

import (
	"database/sql"

	_ "github.com/lib/pq"
)

const (
	connectionString = "host=localhost port=5432 user=postgres dbname=postgres password=1234 sslmode=disable"
)

func ConnectDatabase() *sql.DB {
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		panic(err.Error())
	}
	return db
}
