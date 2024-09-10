package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/KunalDuran/fotuj/internal/data"
)

func (app Application) indexHandler(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("key")

	images, err := app.DB.GetImagesByKey(key)
	if err != nil {
		log.Println(err)
		return
	}

	proj := struct {
		Key    string
		Images []data.Image
	}{
		Key:    key,
		Images: images,
	}

	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		log.Println("Template parsing error:", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, proj)
	if err != nil {
		log.Println("Template execution error:", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}

func (app Application) bucketHandler(w http.ResponseWriter, r *http.Request) {
	var b data.Project

	err := json.NewDecoder(r.Body).Decode(&b)
	if err != nil {
		fmt.Fprint(w, err)
	}

	app.DB.AddProject(b)
	fmt.Fprint(w, "Successfully saved")
}

func (app Application) updateStatusHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		image := r.FormValue("image")
		status := r.FormValue("status")
		key := r.FormValue("key")

		err := app.DB.UpdateImageStatus(key, image, status)
		if err != nil {
			fmt.Fprint(w, "Error in image selection")
			return
		}

		fmt.Fprint(w, "Success")
	}
}
