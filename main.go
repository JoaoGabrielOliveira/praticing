package main

import (
	"fmt"
	"net/http"
)

func main() {
	startApp()
}

func startApp() {
	http.HandleFunc("/", hello)
	http.ListenAndServe(":8090", nil)
}

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello\n")
}
