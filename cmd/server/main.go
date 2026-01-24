package main

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/lib/pq"

	"github.com/Akakazkz/go-task-manager-api/internal/handler"
	"github.com/Akakazkz/go-task-manager-api/internal/repository"
	"github.com/Akakazkz/go-task-manager-api/internal/service"
)

func main() {
	db, err := sql.Open("postgres", "postgres://postgres:kaz.123kz@localhost:5432/taskdb?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	userRepo := repository.NewPostgresUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	mux := http.NewServeMux()

	mux.HandleFunc("/health", handler.Health)
	mux.HandleFunc("/users", userHandler.Create)

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	log.Println("starting server on :8080")
	log.Fatal(server.ListenAndServe())

}
