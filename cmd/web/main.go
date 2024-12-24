package main

import (
	"log"
	"net/http"

	"github.com/jimmydagumjr/GoNAS/internal/handlers"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	// Create a new router
	r := chi.NewRouter()

	// Add middleware
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Define routes
	r.Post("/upload", handlers.FileUploadHandler)

	// Start the server
	log.Println("Starting server on :8080")
	http.ListenAndServe(":8080", r)
}
