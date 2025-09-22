package main

import (
	"html/template"
	"net/http"
)

var tmpl = template.Must(template.ParseFiles("template/index.html"))

func handler(w http.ResponseWriter, r *http.Request) {
	tmpl.Execute(w, nil)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":80", nil)
}
