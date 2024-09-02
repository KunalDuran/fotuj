package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

type ImageStatus struct {
	Image  string `json:"image"`
	Status string `json:"status"` // "selected", "rejected", or ""
}

var images []string
var imageStatuses []ImageStatus

var FILE_PATH = `D:\Kunal Bhaiya Wedding\photo\01 = 29-11-2023\109ND750`
var CONF_PATH = `109ND750`

func loadImages() {
	files, err := os.ReadDir(FILE_PATH)
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		if !file.IsDir() {
			images = append(images, file.Name())
		}
	}

	fmt.Println(images)
}

func loadStatuses() {
	// if does not exist create

	_, err := os.Stat("data/" + CONF_PATH + "status.json")
	if os.IsNotExist(err) {
		file, err := os.Create("data/" + CONF_PATH + "status.json")
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
	}

	// if exist load

	file, err := os.Open("data/" + CONF_PATH + "status.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&imageStatuses); err != nil {
		log.Fatal(err)
	}
}

func saveStatuses() {
	file, err := os.Create("data/" + CONF_PATH + "status.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	encoder := json.NewEncoder(file)
	if err := encoder.Encode(imageStatuses); err != nil {
		log.Fatal(err)
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	data := struct {
		Images        []string
		ImageStatuses []ImageStatus
	}{
		Images:        images,
		ImageStatuses: imageStatuses,
	}
	tmpl.Execute(w, data)
}

func updateStatusHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		image := r.FormValue("image")
		status := r.FormValue("status")

		for i, imgStatus := range imageStatuses {
			if imgStatus.Image == image {
				imageStatuses[i].Status = status
				saveStatuses()
				return
			}
		}
		imageStatuses = append(imageStatuses, ImageStatus{Image: image, Status: status})
		saveStatuses()
	}
}

func main() {
	loadImages()
	loadStatuses()
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/update-status", updateStatusHandler)
	http.Handle("/static/", http.StripPrefix("static", http.FileServer(http.Dir("static"))))
	http.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.Dir(`D:\Kunal Bhaiya Wedding\photo\01 = 29-11-2023\109ND750\`))))
	fmt.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", nil))
}
