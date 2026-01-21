package main

import (
	"log"
	"net/http"

	"github.com/Akakazkz/go-task-manager-api/internal/handler"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/health", handler.Health)

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	log.Println("starting server on :8080")
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
