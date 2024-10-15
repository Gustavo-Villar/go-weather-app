package main

import (
	"log"
	"net/http"

	"github.com/gustavo-villar/go-weather-tracker/internal/api"
)

func main() {
	port := "8080"

	http.HandleFunc("/weather", api.HandleGetWeather)

	log.Printf("Server starting on port %s...\n", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
