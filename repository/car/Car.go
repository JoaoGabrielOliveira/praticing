package car

import (
	"database/sql"
	"praticing/database"
	"praticing/models"
	"praticing/util"
)

func All() []models.Car {
	db := database.ConnectDatabase()

	rows, err := db.Query("SELECT * FROM cars")
	util.CheckError(err)

	return readRow(rows)
}

func Get(id string) models.Car {
	db := database.ConnectDatabase()
	stmt, err := db.Prepare("SELECT * FROM cars WHERE id=$1")
	util.CheckError(err)

	rows, err := stmt.Query(id)
	util.CheckError(err)

	result := readRow(rows)

	if len(result) > 0 {
		return result[0]
	}
	return models.Car{}
}

func readRow(rows *sql.Rows) []models.Car {
	cars := []models.Car{}

	for rows.Next() {
		var id string
		var model_id int
		var price float64

		rows.Scan(&id, &model_id, &price)

		cars = append(cars, models.Car{Id: id, Model_id: model_id, Price: price})
	}
	return cars
}
