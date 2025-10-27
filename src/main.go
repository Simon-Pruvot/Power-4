package main

import "net/http"

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
