package main

import (
	"html/template"
	"net/http"
)

type pageData struct {
	grille [][]int
}

func handler(w http.ResponseWriter, r *http.Request) {
	var tmpl = template.Must(template.ParseFiles("template/index.html"))
	tmpl.Execute(w, nil)
}

func play(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("src/template/play.html"))
	tmpl.Execute(w, nil)
}

func main() {
	http.Handle("/CSS/", http.StripPrefix("/CSS/", http.FileServer(http.Dir("CSS"))))
	http.HandleFunc("/", handler)
	http.HandleFunc("/page2", play)
	http.ListenAndServe(":80", nil)
}
