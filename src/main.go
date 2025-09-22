package main

import (
	"html/template"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	var tmpl = template.Must(template.ParseFiles("template/index.html"))
	tmpl.Execute(w, nil)
}

func main() {
	http.Handle("/CSS/", http.StripPrefix("/CSS/", http.FileServer(http.Dir("CSS"))))
	http.HandleFunc("/", handler)
	http.ListenAndServe(":80", nil)
}
