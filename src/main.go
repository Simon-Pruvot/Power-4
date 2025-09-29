package main

import (
	"html/template"
	"net/http"
)

type pageData struct {
	Grille       [][]string
	joueur       []string
	indiceJoueur int
}

func handler(w http.ResponseWriter, r *http.Request) {
	var tmpl = template.Must(template.ParseFiles("src/template/index.html"))
	tmpl.Execute(w, nil)
}

func play(w http.ResponseWriter, r *http.Request) {
	data := pageData{Grille: [][]string{
		{"/images/pion0.png", "/images/pion0.png", "/images/pion0.png", "/images/pion0.png", "/images/pion0.png", "/images/pion0.png", "/images/pion0.png"},
		{"/images/pion0.png", "/images/pion0.png", "/images/pion0.png", "/images/pion0.png", "/images/pion0.png", "/images/pion0.png", "/images/pion0.png"},
		{"/images/pion0.png", "/images/pion0.png", "/images/pion0.png", "/images/pion0.png", "/images/pion0.png", "/images/pion0.png", "/images/pion0.png"},
		{"/images/pion0.png", "/images/pion0.png", "/images/pion0.png", "/images/pion0.png", "/images/pion0.png", "/images/pion0.png", "/images/pion0.png"},
		{"/images/pion0.png", "/images/pion0.png", "/images/pion0.png", "/images/pion0.png", "/images/pion0.png", "/images/pion0.png", "/images/pion0.png"},
		{"/images/pion0.png", "/images/pion0.png", "/images/pion2.png", "/images/pion1.png", "/images/pion0.png", "/images/pion0.png", "/images/pion0.png"},
	}, joueur: []string{"/images/pion1.png", "/images/pion2.png"}, indiceJoueur: 0}
	tmpl := template.Must(template.ParseFiles("template/play.html"))
	tmpl.Execute(w, data)
}

func (data *pageData) ajouterPion(index int) {
	for i := len(data.Grille); i >= 0; i-- {
		if data.Grille[i][index] == "/images/pion0.png" {
			data.Grille[i][index] = data.joueur[data.indiceJoueur]
		}
	}
}

func main() {
	http.Handle("/CSS/", http.StripPrefix("/CSS/", http.FileServer(http.Dir("CSS"))))
	fs := http.FileServer(http.Dir("./images"))
	http.Handle("/images/", http.StripPrefix("/images/", fs))
	http.HandleFunc("/", handler)
	http.HandleFunc("/play", play)
	http.ListenAndServe(":80", nil)
}

//Palette de couleur:

//: #32346F
//: #4C63C6
//: #FFED19
//: #F5503A
