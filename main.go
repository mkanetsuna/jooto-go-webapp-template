package main

import (
	"log"
	"net/http"
	"os"

	"jooto-go-webapp-template/handlers"
)

func main() {
	http.HandleFunc("/", handlers.HelloHandler)
	http.HandleFunc("/health", handlers.HealthHandler)
	http.HandleFunc("/scheduled-task", handlers.ScheduledTaskHandler)
	http.HandleFunc("/webhook", handlers.WebhookHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Starting server on port %s...", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
