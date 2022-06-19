package models

import "praticing/config"

type Car struct {
	Id       string
	Model_id int
	Price    float64
}

func All(c *Car) []Car {
	db := config.ConnectDatabase()
	cars := []Car{}

	rows, err := db.Query("SELECT * FROM cars")
	config.CheckError(err)
	for rows.Next() {
		var id string
		var model_id int
		var price float64

		rows.Scan(&id, &model_id, &price)

		cars = append(cars, Car{id, model_id, price})
	}
	return cars
}
