package main

import (
	"net/http"
	"text/template"
)

func home(w http.ResponseWriter, r *http.Request) {
	temp := template.Must(template.ParseFiles("template/index.html"))
	temp.Execute(w, nil)
}

func main() {
	http.HandleFunc("/", home)
	http.ListenAndServe(":80", nil)
}
