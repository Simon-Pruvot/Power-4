package main

import (
	"html/template"
	"net/http"
)

var tmpl = template.Must(template.ParseFiles("index.html"))

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":80", nil)
}
