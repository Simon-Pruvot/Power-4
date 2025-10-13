package main

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

type pageData struct {
	Grille       [][]string
	Colonnes     []int
	joueur       []string
	indiceJoueur int
	verifBot     []int
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
	pion1 := r.FormValue("pion1")
	pion2 := r.FormValue("pion2")

	if pion1 != "" {
		data.joueur[0] = "/images/" + pion1 // ✅ fixed path
	}
	if pion2 != "" {
		data.joueur[1] = "/images/" + pion2 // ✅ fixed path
	}

	tmpl := template.Must(template.ParseFiles("template/merch.html", "template/header.html"))

	files, err := os.ReadDir("./images")
	if err != nil {
		http.Error(w, "Unable to read images directory", http.StatusInternalServerError)
		return
	}

	images := []string{}
	for _, f := range files {
		if f.IsDir() {
			continue
		}
		images = append(images, f.Name())
	}

	// ✅ Don’t skip the first image
	tmpl.Execute(w, images)
}

func pers(w http.ResponseWriter, r *http.Request) {
	var tmpl = template.Must(template.ParseFiles("template/personalisation.html", "template/header.html"))
	tmpl.Execute(w, nil)
}

func play(w http.ResponseWriter, r *http.Request) {
	rowsStr := r.FormValue("rows")
	colsStr := r.FormValue("cols")

	rows, cols := 6, 7
	if rInt, err := strconv.Atoi(rowsStr); err == nil {
		rows = rInt
	}
	if cInt, err := strconv.Atoi(colsStr); err == nil {
		cols = cInt
	}

	// Always recreate the grid if rows/cols are present in the request
	if rowsStr != "" && colsStr != "" && rows != 6 && cols != 7 {
		data.Grille = nouvelleGrille(rows, cols)
		data.joueur = []string{"/images/pion1.png", "/images/pion2.png"}
		data.indiceJoueur = 0
		data.Colonnes = make([]int, cols)
	}

	// Handle column click
	colStr := r.FormValue("col")
	if colStr != "" {
		col, err := strconv.Atoi(colStr)
		if err == nil {
			if data.ajouterPion(col) {
				http.Redirect(w, r, "/", http.StatusSeeOther)
				return
			}
		}
	}

	for i := range data.Colonnes {
		data.Colonnes[i] = i
	}

	// Compute number of columns from the actual grid
	numCols := 0
	if len(data.Grille) > 0 {
		numCols = len(data.Grille[0])
	}

	// Render template
	tmpl := template.Must(template.ParseFiles("template/play.html", "template/header.html"))
	tmpl.Execute(w, struct {
		pageData
		ColsPlusOne int
	}{data, numCols})

}

func (data *pageData) verif(ligne int, col int) int {
	max := 1
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
	if max < compteur {
		max = compteur
	}
	if compteur >= 4 {
		return max
	} else {
		compteur = 1
		for i := col; i < len(data.Grille)-1; i++ {
			if data.Grille[i][ligne] == data.Grille[col][ligne] {
				compteur += 1
			} else {
				break
			}
		}
		if max < compteur {
			max = compteur
		}
		if compteur >= 3 {
			return compteur
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
			if max < compteur {
				max = compteur
			}
			if compteur >= 4 {
				return compteur
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
				if max < compteur {
					max = compteur
				}
				if compteur >= 4 {
					return compteur
				}
			}

		}
	}
	return max
}

func (data *pageData) ajouterPion(index int) bool {
	for i := len(data.Grille) - 1; i >= 0; i-- {
		if data.Grille[i][index] == "/images/pion0.png" {
			data.Grille[i][index] = data.joueur[data.indiceJoueur]
			if data.verif(index, i) >= 4 {
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
	data.verifBot = make([]int, len(data.Grille[0]))
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

// GET /camera — show the webcam capture page
func cameraPage(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("template/camera.html", "template/header.html"))
	tmpl.Execute(w, nil)
}

func uploadPhoto(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	file, _, err := r.FormFile("photo")
	if err != nil {
		http.Error(w, "Missing photo", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// generate a random name
	name := fmt.Sprintf("photo_%d.jpg", time.Now().UnixNano())
	outPath := filepath.Join("images", name)

	out, err := os.Create(outPath)
	if err != nil {
		http.Error(w, "Failed to create file", http.StatusInternalServerError)
		return
	}
	defer out.Close()

	_, err = io.Copy(out, file)
	if err != nil {
		http.Error(w, "Failed to save file", http.StatusInternalServerError)
		return
	}

	log.Println("Saved new photo:", outPath)

	http.Redirect(w, r, "/merch", http.StatusSeeOther)
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
