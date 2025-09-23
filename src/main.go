package main

import (
	"html/template"
	"net/http"
)

type pageData struct {
	grille [][]int
}
jeu:= pageData{{{0,0,0,0,0,0,0},
{0,0,0,0,0,0,0},
{0,0,0,0,0,0,0},
{0,0,0,0,0,0,0},
{0,0,0,0,0,0,0},
{0,0,0,0,0,0,0},}}
func handler(w http.ResponseWriter, r *http.Request) {
	var tmpl = template.Must(template.ParseFiles("template/index.html"))
	tmpl.Execute(w, nil)
}

func play(w http.ResponseWriter, r *http.Request) {
	var tmpl = template.Must(template.ParseFiles("template/play.html"))
	tmpl.Execute(w, nil)
}

func main() {
	http.Handle("/CSS/", http.StripPrefix("/CSS/", http.FileServer(http.Dir("CSS"))))
	fs := http.FileServer(http.Dir("./images"))
	http.Handle("/images/", http.StripPrefix("/images/", fs))
	http.HandleFunc("/", handler)
	http.HandleFunc("/play", play)
	http.ListenAndServe(":80", nil)
}
