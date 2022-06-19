package config

import (
	"html/template"
	"net/http"
)

var temp = template.Must(template.ParseGlob("views/*.html"))

func StartApp() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	err := temp.ExecuteTemplate(w, "Index", nil)

	if err != nil {
		panic(err.Error())
	}
}
