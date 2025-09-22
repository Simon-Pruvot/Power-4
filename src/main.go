package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var tmpl = template.Must(template.ParseFiles("index.html"))

func handler(w http.ResponseWritter, r *http.Request) {
	fmt.Fprintf(w, "Hello World !")
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":80", nil)
}
