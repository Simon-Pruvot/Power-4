package main

import (
	"html/template"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("template/index.html", "template/header.html"))
	tmpl.Execute(w, nil)
}

func diff(w http.ResponseWriter, r *http.Request) {
	var tmpl = template.Must(template.ParseFiles("template/diff.html", "template/header.html"))
	tmpl.Execute(w, nil)
}

func victoire(w http.ResponseWriter, r *http.Request) {
	joueur := r.URL.Query().Get("winner")
	var tmpl = template.Must(template.ParseFiles("template/victoire.html", "template/header.html"))
	tmpl.Execute(w, joueur)
}

func egalite(w http.ResponseWriter, r *http.Request) {
	var tmpl = template.Must(template.ParseFiles("template/egalite.html", "template/header.html"))
	tmpl.Execute(w, egalite)
}

func regle(w http.ResponseWriter, r *http.Request) {
	var tmpl = template.Must(template.ParseFiles("template/regle.html", "template/header.html"))
	tmpl.Execute(w, regle)
}

func pers(w http.ResponseWriter, r *http.Request) {
	var tmpl = template.Must(template.ParseFiles("template/personalisation.html", "template/header.html"))
	tmpl.Execute(w, nil)
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
	http.HandleFunc("/camera", cameraPage)
	http.HandleFunc("/uploadphoto", uploadPhoto)
	http.HandleFunc("/victoire", victoire)
	http.HandleFunc("/egalite", egalite)
	http.HandleFunc("/regle", regle)
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
