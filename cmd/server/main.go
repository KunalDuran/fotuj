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
	http.HandleFunc("/bucket", app.bucketHandler)
	http.HandleFunc("/select", app.updateStatusHandler)

	fs := http.FileServer(http.Dir("./output"))
	http.Handle("/images/", http.StripPrefix("/images/", fs))

	fmt.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", nil))
}
