package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/status", HandleStatus)

	log.Print("Starting server")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func HandleStatus(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
