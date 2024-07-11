package main

import (
	"fmt"
	handlers "go_race_condition/src/handlers"
	repo "go_race_condition/src/repository"
	"go_race_condition/src/router"
	usecases "go_race_condition/src/usecases"
	"log"
	"net/http"
)

func main() {
	// Initialize the application by creating instances of repository, usecase, handler, and router.
	// The product repository is initialized with an initial stock of 100000.
	productRepo := repo.NewProductRepository(100000)
	productUsecase := usecases.NewProductUsecase(productRepo)
	handler := handlers.NewHandler(productUsecase)

	// Create a new router instance using the handler to handle incoming HTTP requests.
	r := router.NewRouter(handler)

	// Print a message indicating that the server is running on port 8080 and start the HTTP server.
	fmt.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
