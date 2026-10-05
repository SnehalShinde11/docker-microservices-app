package main

import (
	"fmt"
	"log"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")

	fmt.Fprintln(w, "Backend Service is running - Webhook Test - Jenkins Webhook Test")
	fmt.Fprintln(w, "This application is designed by Snehal for Demo Purpose")
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintln(w, "Backend Service is healthy")
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/health", healthHandler)

	log.Println("========================================")
	log.Println("Snehal Microservices Backend")
	log.Println("Backend Service running on port 8080")
	log.Println("========================================")

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Fatal(err)
	}
}

