package main

import (
	"log"
	"net/http"

	"github.com/Akakazkz/go-task-manager-api/internal/handler"
	"github.com/Akakazkz/go-task-manager-api/internal/service"
)

func main() {
	mux := http.NewServeMux()

	userService := service.NewUserService()
	userHandler := handler.NewUserHandler(userService)

	mux.HandleFunc("/health", handler.Health)
	mux.HandleFunc("/users", userHandler.Create)

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	log.Println("starting server on :8080")
	log.Fatal(server.ListenAndServe())
}
