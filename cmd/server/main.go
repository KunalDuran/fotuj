package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/KunalDuran/fotuj/internal/data"
)

type Application struct {
	DB data.Database
}

func NewApplication() *Application {
	return &Application{
		DB: data.NewSQLiteDB(),
	}
}

func main() {
	app := NewApplication()
	http.HandleFunc("/", app.indexHandler)
	http.HandleFunc("/gallery", app.galleryHandler)
	http.HandleFunc("/select", app.updateStatusHandler)

	fs := http.FileServer(http.Dir("./output"))
	http.Handle("/images/", http.StripPrefix("/images/", fs))
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	fmt.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", nil))
}
