package errors

import (
	"database/sql"
)

func CheckDatabaseConnection(connection *sql.DB) error {
	if connection.Ping() != nil {
		return ERR_DATABASE_CONNECTION(*connection)
	}
	return nil
}
