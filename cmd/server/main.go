package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

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
	}

	fmt.Println(os.Getwd())

	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		fmt.Errorf("Template not found")
	}

	tmpl.Execute(w, result)
}

func bucketHandler(w http.ResponseWriter, r *http.Request) {
	var b data.Bucket
	if err := json.NewDecoder(r.Body).Decode(&b); err != nil {
		log.Println(err)
		fmt.Fprint(w, err)
	}

	data.InsertOne(data.COLLECTION_BUCKET, b)
	fmt.Fprint(w, "Successfully saved")
}

func updateStatusHandler(w http.ResponseWriter, r *http.Request) {
	// body, err := io.ReadAll(r.Body)
	// if err != nil {
	// 	log.Println(err)
	// }

}

func main() {

	data.InitDB("")

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/bucket", bucketHandler)
	http.HandleFunc("/update-status", updateStatusHandler)

	fs := http.FileServer(http.Dir("./output"))
	http.Handle("/images/", http.StripPrefix("/images/", fs))

	fmt.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", nil))
}
