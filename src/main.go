package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strconv"
)

type pageData struct {
	Grille       [][]string
	Colonnes     []int
	joueur       []string
	indiceJoueur int
}

var data pageData

func handler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("template/index.html", "template/header.html"))
	tmpl.Execute(w, nil)
}

func diff(w http.ResponseWriter, r *http.Request) {
	var tmpl = template.Must(template.ParseFiles("template/diff.html", "template/header.html"))
	tmpl.Execute(w, nil)
}

func merch(w http.ResponseWriter, r *http.Request) {
	pion := r.FormValue("pion")
	fmt.Println("a: ", pion)
	if pion != "" {
		data.joueur[0] = "../images/" + pion
	}
	var tmpl = template.Must(template.ParseFiles("template/merch.html", "template/header.html"))
	temp, _ := os.ReadDir("./images")
	images := []string{}
	for _, e := range temp {
		images = append(images, e.Name())
	}
	tmpl.Execute(w, images[1:])
}

func pers(w http.ResponseWriter, r *http.Request) {
	var tmpl = template.Must(template.ParseFiles("template/personalisation.html", "template/header.html"))
	tmpl.Execute(w, nil)
}

func play(w http.ResponseWriter, r *http.Request) {
	colStr := r.FormValue("col")
	if colStr != "" {
		col, err := strconv.Atoi(colStr)
		if err == nil {
			if data.ajouterPion(col) {
				http.Redirect(w, r, "/", http.StatusSeeOther)
			}
		}
	}
	for index := range data.Colonnes {
		data.Colonnes[index] = index
	}
	tmpl := template.Must(template.ParseFiles("template/play.html", "template/header.html"))
	tmpl.Execute(w, data)
}
func (data *pageData) verif(ligne int, col int) bool {
	compteur := 1
	for i := ligne - 1; i > -1; i-- {
		if data.Grille[col][i] == data.Grille[col][ligne] {
			compteur += 1
		} else {
			break
		}
	}
	for i := ligne + 1; i < len(data.Grille[0])-1; i++ {
		if data.Grille[col][i] == data.Grille[col][ligne] {
			compteur += 1
		} else {
			break
		}
	}
	if compteur >= 4 {
		return true
	} else {
		compteur = 1
		for i := col + 1; i < len(data.Grille)-1; i++ {
			if data.Grille[i][ligne] == data.Grille[col][ligne] {
				compteur += 1
			} else {
				break
			}
		}
		if compteur >= 3 {
			return true
		} else {
			compteur = 1
			for i := 1; col-i >= 0 && ligne-i >= 0; i++ {
				if data.Grille[col-i][ligne-i] == data.Grille[col][ligne] {
					compteur += 1
				} else {
					break
				}
			}
			for i := 1; col+i <= len(data.Grille)-1 && ligne+i <= len(data.Grille[0])-1; i++ {
				if data.Grille[col+i][ligne+i] == data.Grille[col][ligne] {
					compteur += 1
				} else {
					break
				}
			}
			if compteur >= 4 {
				return true
			} else {
				compteur = 1
				for i := 1; col+i <= len(data.Grille)-1 && ligne-i >= 0; i++ {
					if data.Grille[col+i][ligne-i] == data.Grille[col][ligne] {
						compteur += 1
					} else {
						break
					}
				}
				for i := 1; col-i >= 0 && ligne+i <= len(data.Grille[0])-1; i++ {
					if data.Grille[col-i][ligne+i] == data.Grille[col][ligne] {
						compteur += 1
					} else {
						break
					}
				}
				if compteur >= 4 {
					return true
				}
			}

		}
	}
	return false
}

func (data *pageData) ajouterPion(index int) bool {
	for i := len(data.Grille) - 1; i >= 0; i-- {
		if data.Grille[i][index] == "/images/pion0.png" {
			data.Grille[i][index] = data.joueur[data.indiceJoueur]
			if data.verif(index, i) {
				return true
			}
			data.indiceJoueur = (data.indiceJoueur + 1) % 2
			break
		}
	}
	return false
}

func temp(w http.ResponseWriter, r *http.Request) {

	data.Grille = [][]string{
		{"/images/pion0.png", "/images/pion0.png", "/images/pion0.png", "/images/pion0.png", "/images/pion0.png", "/images/pion0.png", "/images/pion0.png"},
		{"/images/pion0.png", "/images/pion0.png", "/images/pion0.png", "/images/pion0.png", "/images/pion0.png", "/images/pion0.png", "/images/pion0.png"},
		{"/images/pion0.png", "/images/pion0.png", "/images/pion0.png", "/images/pion0.png", "/images/pion0.png", "/images/pion0.png", "/images/pion0.png"},
		{"/images/pion0.png", "/images/pion0.png", "/images/pion0.png", "/images/pion0.png", "/images/pion0.png", "/images/pion0.png", "/images/pion0.png"},
		{"/images/pion0.png", "/images/pion0.png", "/images/pion0.png", "/images/pion0.png", "/images/pion0.png", "/images/pion0.png", "/images/pion0.png"},
		{"/images/pion0.png", "/images/pion0.png", "/images/pion0.png", "/images/pion0.png", "/images/pion0.png", "/images/pion0.png", "/images/pion0.png"},
	}
	data.Colonnes = make([]int, len(data.Grille[0]))
	fmt.Println(data)
	http.Redirect(w, r, "/play", http.StatusSeeOther)
}

func nouvelleGrille(rows, cols int) [][]string {
	grille := make([][]string, rows)
	for i := range grille {
		grille[i] = make([]string, cols)
		for j := range grille[i] {
			grille[i][j] = "/images/pion0.png"
		}
	}
	return grille
}

func main() {
	data = pageData{joueur: []string{"/images/pion1.png", "/images/pion2.png"}, indiceJoueur: 0}
	http.Handle("/CSS/", http.StripPrefix("/CSS/", http.FileServer(http.Dir("CSS"))))
	fs := http.FileServer(http.Dir("./images"))
	http.Handle("/images/", http.StripPrefix("/images/", fs))
	http.HandleFunc("/", handler)
	http.HandleFunc("/play", play)
	http.HandleFunc("/diff", diff)
	http.HandleFunc("/temp", temp)
	http.HandleFunc("/merch", merch)

	http.HandleFunc("/personalisation", pers)
	http.ListenAndServe(":80", nil)
}

//Palette de couleur:

//: #32346F
//: #4C63C6
//: #FFED19
//: #F5503A

/*func nouvelleGrille(rows, cols int) [][]string {
	grille := make([][]string, rows)
	for i := range grille {
		grille[i] = make([]string, cols)
		for j := range grille[i] {
			grille[i][j] = "/images/pion0.png"
		}
	}
	return grille
}*/

/*rowsStr := r.FormValue("rows")
colsStr := r.FormValue("cols")

rows, cols := 6, 7 // valeurs par défaut
if rInt, err := strconv.Atoi(rowsStr); err == nil {
	rows = rInt
}
if cInt, err := strconv.Atoi(colsStr); err == nil {
	cols = cInt
}

// Créer la grille dynamique
data = pageData{
	Grille:       nouvelleGrille(rows, cols),
	joueur:       []string{"/images/pion1.png", "/images/pion2.png"},
	indiceJoueur: 0,
}*/
