package main

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"slices"
	"strconv"
	"strings"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("src/template/index.html", "src/template/header.html"))
	tmpl.Execute(w, nil)
}

func diff(w http.ResponseWriter, r *http.Request) {
	var tmpl = template.Must(template.ParseFiles("src/template/diff.html", "src/template/header.html"))
	tmpl.Execute(w, nil)
}

func victoire(w http.ResponseWriter, r *http.Request) {
	joueur := r.URL.Query().Get("winner")
	var tmpl = template.Must(template.ParseFiles("src/template/victoire.html", "src/template/header.html"))
	tmpl.Execute(w, joueur)
}

func egalite(w http.ResponseWriter, r *http.Request) {
	var tmpl = template.Must(template.ParseFiles("src/template/egalite.html", "src/template/header.html"))
	tmpl.Execute(w, egalite)
}

func regle(w http.ResponseWriter, r *http.Request) {
	var tmpl = template.Must(template.ParseFiles("src/template/regle.html", "src/template/header.html"))
	tmpl.Execute(w, regle)
}

func merch(w http.ResponseWriter, r *http.Request) {
	id, data := visitor(w, r)

	// Jetons proposés : les images du jeu + les photos prises par CE visiteur.
	files, err := os.ReadDir("src/images")
	if err != nil {
		http.Error(w, "Unable to read images directory", http.StatusInternalServerError)
		return
	}

	images := []string{}
	for _, f := range files {
		if f.IsDir() {
			continue
		}
		images = append(images, "/images/"+f.Name())
	}
	if uploads, err := os.ReadDir(uploadDir()); err == nil {
		for _, f := range uploads {
			if !f.IsDir() && strings.HasPrefix(f.Name(), id+"_") {
				images = append(images, "/uploads/"+f.Name())
			}
		}
	}

	// On n'accepte qu'un jeton de la liste ci-dessus.
	if pion1 := r.FormValue("pion1"); slices.Contains(images, pion1) {
		data.joueur[0] = pion1
	}
	if pion2 := r.FormValue("pion2"); slices.Contains(images, pion2) {
		data.joueur[1] = pion2
	}

	tmpl := template.Must(template.ParseFiles("src/template/merch.html", "src/template/header.html"))
	tmpl.Execute(w, images)
}

func pers(w http.ResponseWriter, r *http.Request) {
	var tmpl = template.Must(template.ParseFiles("src/template/personalisation.html", "src/template/header.html"))
	tmpl.Execute(w, nil)
}

func play(w http.ResponseWriter, r *http.Request) {
	data := gameFor(w, r)
	rowsStr := r.FormValue("rows")
	colsStr := r.FormValue("cols")

	rows, cols := 6, 7
	if rInt, err := strconv.Atoi(rowsStr); err == nil {
		rows = rInt
	}
	if cInt, err := strconv.Atoi(colsStr); err == nil {
		cols = cInt
	}

	if rowsStr != "" && colsStr != "" {
		data.Grille = nouvelleGrille(rows, cols)
		data.joueur = []string{"/images/pion1.png", "/images/pion2.png"}
		data.indiceJoueur = 0
		data.Colonnes = make([]int, cols)
	}

	colStr := r.FormValue("col")
	if colStr != "" {
		col, err := strconv.Atoi(colStr)
		if err == nil {
			if data.ajouterPion(col, data.Grille, w, r) {
				http.Redirect(w, r, "/victoire?winner="+data.joueur[data.indiceJoueur], http.StatusSeeOther)
				return
			}
		}
	}

	for i := range data.Colonnes {
		data.Colonnes[i] = i
	}

	numCols := 0
	if len(data.Grille) > 0 {
		numCols = len(data.Grille[0])
	}

	tmpl := template.Must(template.ParseFiles("src/template/play.html", "src/template/header.html"))
	tmpl.Execute(w, struct {
		pageData
		ColsPlusOne int
	}{*data, numCols})

}

/*
func RechercheBot(profondeur int, grille [][]string, temp [][]int) {
	copie := make([][]string, len(grille))
	for i := range grille {
		copie[i] = append([]string(nil), grille[i]...)
	}
	for i := 0; i < len(grille[0]); i++ {
		ligne := chercherLigne(i)
		if ligne != -1 {
			verif := data.verif(i, ligne, data.joueur[profondeur%2])
			if verif > temp[i][1] && profondeur%2 == 0 {
				temp[profondeur%2][1] = i
				temp[profondeur%2][0] = verif
			} else if verif < temp[i][1] && profondeur%2 == 1 {
				temp[profondeur%2][1] = i
				temp[profondeur%2][0] = verif
			}
		}
	}
	if profondeur%2 < len(temp) && len(temp[profondeur%2]) > 1 {
		ligne := chercherLigne(temp[profondeur%2][1])
		if ligne != -1 {
			copie[ligne][temp[profondeur%2][1]] = data.joueur[profondeur%2]
		}
	}
	if profondeur == data.botProfondeur {
		fmt.Println("choix du bot:", temp)
		return
	} else {
		RechercheBot(profondeur+1, data.Grille, temp)
	}
}
*/

func temp(w http.ResponseWriter, r *http.Request) {
	data := gameFor(w, r)
	data.Grille = [][]string{
		{"/images/pion0.png", "/images/pion0.png", "/images/pion0.png", "/images/pion0.png", "/images/pion0.png", "/images/pion0.png", "/images/pion0.png"},
		{"/images/pion0.png", "/images/pion0.png", "/images/pion0.png", "/images/pion0.png", "/images/pion0.png", "/images/pion0.png", "/images/pion0.png"},
		{"/images/pion0.png", "/images/pion0.png", "/images/pion0.png", "/images/pion0.png", "/images/pion0.png", "/images/pion0.png", "/images/pion0.png"},
		{"/images/pion0.png", "/images/pion0.png", "/images/pion0.png", "/images/pion0.png", "/images/pion0.png", "/images/pion0.png", "/images/pion0.png"},
		{"/images/pion0.png", "/images/pion0.png", "/images/pion0.png", "/images/pion0.png", "/images/pion0.png", "/images/pion0.png", "/images/pion0.png"},
		{"/images/pion0.png", "/images/pion0.png", "/images/pion0.png", "/images/pion0.png", "/images/pion0.png", "/images/pion0.png", "/images/pion0.png"},
	}
	data.Colonnes = make([]int, len(data.Grille[0]))
	http.Redirect(w, r, "/play", http.StatusSeeOther)
}

func cameraPage(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("src/template/camera.html", "src/template/header.html"))
	tmpl.Execute(w, nil)
}

func uploadPhoto(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	id, _ := visitor(w, r)
	r.Body = http.MaxBytesReader(w, r.Body, 5<<20) // 5 Mo max par photo

	file, _, err := r.FormFile("photo")
	if err != nil {
		http.Error(w, "Missing photo", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// La photo n'est visible que par le visiteur qui l'a prise (préfixe = son identifiant).
	if err := os.MkdirAll(uploadDir(), 0o755); err != nil {
		http.Error(w, "Failed to create upload directory", http.StatusInternalServerError)
		return
	}
	name := fmt.Sprintf("%s_%d.jpg", id, time.Now().UnixNano())
	outPath := filepath.Join(uploadDir(), name)

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

// uploadDir : dossier des photos prises avec la webcam (UPLOAD_DIR, « uploads » par défaut).
func uploadDir() string {
	if dir := os.Getenv("UPLOAD_DIR"); dir != "" {
		return dir
	}
	return "uploads"
}
