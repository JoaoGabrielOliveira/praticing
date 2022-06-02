package router

import (
	"html/template"
	"net/http"
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
	err := temp.ExecuteTemplate(w, "Index", car.All())

	util.CheckError(err)
}
