package main

import (
	"fmt"
	"log"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {

}
func updateStatusHandler(w http.ResponseWriter, r *http.Request) {

}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/update-status", updateStatusHandler)

	fmt.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", nil))
}
