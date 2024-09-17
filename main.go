package main

import (
	"html/template"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)

	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/index.html"))

	data := struct {
		Title string
	}{
		Title: "Fotuj",
	}
	tmpl.Execute(w, data)
}
