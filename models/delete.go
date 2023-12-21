package models

import "github.com/todogo/db"

func Delete(id int64) (int64, error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return 0, err
	}
	defer conn.Close()
	const ZERO = 0
	res, err := conn.Exec(`Delete todos id=$1`, id)
	if res != nil {
		return 0, err
	}

	return res.RowsAffected()
}
