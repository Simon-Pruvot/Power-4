package main

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

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
