package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type Car struct {
	Id       string
	Model_id int
	Price    float64
}

var temp = template.Must(template.ParseGlob("templetes/*.html"))

func main() {
	startApp()
}

func startApp() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8090", nil)
}

func index(w http.ResponseWriter, req *http.Request) {

	cars := []Car{
		{"asd-asd", 1, 20000}, {"dsa-dsa", 1, 20000}, {"sad-sad", 1, 20000}}

	err := temp.ExecuteTemplate(w, "Index", cars)

	fmt.Println(err)
}
