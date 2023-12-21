package models

import "github.com/todogo/db"

func Update(todo Todo) (int64, error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return 0, err
	}
	defer conn.Close()
	const ZERO = 0
	res, err := conn.Exec(`UPDATE todos SET title=$2, description=$3, done=$4 WHERE id=$1`, todo.ID, todo.Title, todo.Description, todo.Done)
	if res != nil {
		return 0, err
	}

	return res.RowsAffected()
}
