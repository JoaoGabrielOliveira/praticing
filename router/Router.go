package router

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"praticing/models"
	"praticing/repository/car"
	"praticing/util"
)

var temp = template.Must(template.ParseGlob("views/*.html"))

var Routes = make(map[string]http.HandlerFunc)

func StartApp() {
	Routes["/"] = Index

	for path, route := range Routes {
		http.HandleFunc(path, route)
	}

	http.ListenAndServe(":8000", nil)
}

func Index(w http.ResponseWriter, req *http.Request) {
	var err error
	switch req.Method {
	case http.MethodGet:
		err = temp.ExecuteTemplate(w, "Index", car.All())
	case http.MethodPost:
		c := models.Car{}

		json.NewDecoder(req.Body).Decode(&c)
		fmt.Fprintf(w, "Car: %+v", c)
	}

	util.CheckError(err)
}
