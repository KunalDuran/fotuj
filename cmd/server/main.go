package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/KunalDuran/image-reviewer/internal/data"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("key")

	filter := map[string]interface{}{
		"key": key,
	}

	var result data.Bucket
	err := data.FindOne(data.COLLECTION_BUCKET, filter, &result)
	if err != nil {
		log.Println(err)
		return
	}

	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		log.Println("Template parsing error:", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, result)
	if err != nil {
		log.Println("Template execution error:", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}

func bucketHandler(w http.ResponseWriter, r *http.Request) {
	var b data.Bucket

	err := json.NewDecoder(r.Body).Decode(&b)
	if err != nil {
		fmt.Fprint(w, err)
	}

	data.InsertOne(data.COLLECTION_BUCKET, b)
	fmt.Fprint(w, "Successfully saved")
}

func updateStatusHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		image := r.FormValue("image")
		status := r.FormValue("status")
		key := r.FormValue("key")

		fmt.Println(image, status, key)

		statusInt, err := strconv.Atoi(status)
		if err != nil {
			fmt.Fprint(w, "Wrong status")
			return
		}

		err = data.UpdateStatus(key, image, statusInt)
		if err != nil {
			fmt.Fprint(w, "Error in image selection")
			return
		}

		fmt.Fprint(w, "Success")
	}
}

func main() {

	data.InitDB("")

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/bucket", bucketHandler)
	http.HandleFunc("/select", updateStatusHandler)

	fs := http.FileServer(http.Dir("./output"))
	http.Handle("/images/", http.StripPrefix("/images/", fs))

	fmt.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", nil))
}
