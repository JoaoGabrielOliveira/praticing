package router

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"praticing/models"
	CarRepository "praticing/repository/car"
	"praticing/util"
)

var temp = template.Must(template.ParseGlob("views/*.html"))

var Routes = map[string]http.HandlerFunc{
	"/": Index, "/car": Car}

func StartApp() {

	for path, route := range Routes {
		http.HandleFunc(path, route)
	}

	http.ListenAndServe(":8000", nil)
}

func Index(w http.ResponseWriter, req *http.Request) {
	var err error
	switch req.Method {
	case http.MethodGet:
		err = temp.ExecuteTemplate(w, "Index", CarRepository.All())
	}

	util.CheckError(err)
}

func Car(w http.ResponseWriter, req *http.Request) {
	var err error
	switch req.Method {
	case http.MethodGet:
		id, ok := req.URL.Query()["id"]
		util.CheckError(err)

		if !ok || len(id[0]) < 1 {
			fmt.Println("Url Param 'key' is missing")
			return
		} else {
			err = temp.ExecuteTemplate(w, "Car", CarRepository.Get(id[0]))
		}

	case http.MethodPost:
		c := models.Car{}

		err = json.NewDecoder(req.Body).Decode(&c)
		fmt.Fprintf(w, "Car: %+v", c)
	}

	util.CheckError(err)
}
