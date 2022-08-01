package errors

import (
	"database/sql"
	"fmt"
)

const (
	ERR_600 = "Can't connect to database"
)

type ErrorMessege struct {
	id      string
	message string
	data    any
}

func (err ErrorMessege) Error() string {
	return fmt.Sprintf("%v - %v %+v", err.id, err.message, err.data)
}

func ERR_DATABASE_CONNECTION(conn sql.DB) ErrorMessege {
	return ErrorMessege{"600", ERR_600, conn}
}
