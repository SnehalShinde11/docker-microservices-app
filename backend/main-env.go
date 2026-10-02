package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)

	if value == "" {
		return defaultValue
	}

	return value
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")

	appName := getEnv("APP_NAME", "Snehal Microservices Backend")
	appEnvironment := getEnv("APP_ENV", "development")
	dbHost := getEnv("DB_HOST", "localhost")
	dbPort := getEnv("DB_PORT", "5432")
	dbName := getEnv("DB_NAME", "microservices")
	dbUser := getEnv("DB_USER", "snehal")

	fmt.Fprintln(w, "Backend Service is running")
	fmt.Fprintln(w, "Application:", appName)
	fmt.Fprintln(w, "Environment:", appEnvironment)
	fmt.Fprintln(w, "Database Host:", dbHost)
	fmt.Fprintln(w, "Database Port:", dbPort)
	fmt.Fprintln(w, "Database Name:", dbName)
	fmt.Fprintln(w, "Database User:", dbUser)
	fmt.Fprintln(w, "This application is designed by Snehal for Demo Purpose")
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")

	fmt.Fprintln(w, "Backend Service is healthy")
}

func main() {
	port := getEnv("PORT", "8080")
	appName := getEnv("APP_NAME", "Snehal Microservices Backend")

	log.Println("========================================")
	log.Println(appName)
	log.Println("Backend Service running on port", port)
	log.Println("========================================")

	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/health", healthHandler)

	err := http.ListenAndServe(":"+port, nil)

	if err != nil {
		log.Fatal(err)
	}
}


