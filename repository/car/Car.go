package car

import (
	"praticing/database"
	"praticing/models"
	"praticing/util"
)

func All() []models.Car {
	db := database.ConnectDatabase()
	cars := []models.Car{}

	rows, err := db.Query("SELECT * FROM cars")
	util.CheckError(err)
	for rows.Next() {
		var id string
		var model_id int
		var price float64

		rows.Scan(&id, &model_id, &price)

		cars = append(cars, models.Car{id, model_id, price})
	}
	return cars
}
