package main

import (
	"html/template"
	"net/http"
	"os"
)

func merch(w http.ResponseWriter, r *http.Request) {
	pion1 := r.FormValue("pion1")
	pion2 := r.FormValue("pion2")

	if pion1 != "" {
		data.joueur[0] = "/images/" + pion1 // fixed path
	}
	if pion2 != "" {
		data.joueur[1] = "/images/" + pion2 // fixed path
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

	// Don’t skip the first image
	tmpl.Execute(w, images)
}
