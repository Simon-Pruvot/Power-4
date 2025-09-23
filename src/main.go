package main

import (
	"html/template"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	var tmpl = template.Must(template.ParseFiles("src/template/index.html"))
	tmpl.Execute(w, nil)
}

func main() {
	http.Handle("/CSS/", http.StripPrefix("/CSS/", http.FileServer(http.Dir("src/CSS"))))
	http.HandleFunc("/", handler)
	http.ListenAndServe(":80", nil)
}
