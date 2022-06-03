package car

import (
	"database/sql"
	"praticing/database"
	"praticing/models"
	"praticing/util"
)

func All() ([]models.Car, error) {
	db, err := database.ConnectDatabase()
	if err != nil {
		return nil, err
	}

	rows, err := db.Query("SELECT * FROM cars")
	if util.CheckError(err, nil) {
		return nil, err
	}

	result, err := readRow(rows)
	if util.CheckError(err, nil) {
		return nil, err
	}

	return result, nil
}

func Get(id string) models.Car {
	db, err := database.ConnectDatabase()
	if util.CheckError(err, nil) {
		return models.Car{}
	}

	stmt, err := db.Prepare("SELECT * FROM cars WHERE id=$1")
	util.CheckError(err, nil)

	rows, err := stmt.Query(id)
	util.CheckError(err, nil)
	result, err := readRow(rows)
	util.CheckError(err, nil)

	if len(result) > 0 {
		return result[0]
	}
	return models.Car{}
}

func readRow(rows *sql.Rows) ([]models.Car, error) {
	cars := []models.Car{}

	for rows.Next() {
		var id string
		var model_id int
		var price float64

		err := rows.Scan(&id, &model_id, &price)
		if util.CheckError(err, nil) {
			return nil, err
		}

		cars = append(cars, models.Car{Id: id, Model_id: model_id, Price: price})
	}
	return cars, nil
}
