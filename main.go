package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"

	_ "github.com/lib/pq"
)

type Car struct {
	Id       string
	Model_id int
	Price    float64
}

var temp = template.Must(template.ParseGlob("views/*.html"))

func main() {
	startApp()
}

func startApp() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	db := connect_database("host=localhost port=5432 user=postgres dbname=postgres password=1234 sslmode=disable")

	car := Car{}
	cars := []Car{}

	cars_rows, err := db.Query("SELECT * FROM cars")
	for cars_rows.Next() {
		var Id string
		var Model_id int
		var Price float64

		row_error := cars_rows.Scan(&Id, &Model_id, &Price)
		if row_error != nil {
			panic(row_error.Error())
		}

		car.Id = Id
		car.Model_id = Model_id
		car.Price = Price

		cars = append(cars, car)
	}
	err = temp.ExecuteTemplate(w, "Index", cars)
	defer db.Close()

	fmt.Println(err)
}

func connect_database(conn string) *sql.DB {
	db, err := sql.Open("postgres", conn)
	if err != nil {
		panic(err.Error())
	}
	return db
}
