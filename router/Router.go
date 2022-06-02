package router

import (
	"html/template"
	"net/http"
	"praticing/models"
)

var temp = template.Must(template.ParseGlob("../views/*.html"))

var Routes = make(map[string]http.HandlerFunc)

func StartApp() {
	Routes["/"] = Index

	for path, route := range Routes {
		http.HandleFunc(path, route)
	}
	//http.HandleFunc("/", Index)

	http.ListenAndServe(":8000", nil)
}

func Index(w http.ResponseWriter, req *http.Request) {
	cars := models.All()

	err := temp.ExecuteTemplate(w, "Index", cars)

	if err != nil {
		panic(err.Error())
	}
}
