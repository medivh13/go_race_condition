package main

import (
    "fmt"
    "log"
    "net/http"
    handlers "go_race_condition/src/handlers"
    repo "go_race_condition/src/repository"
    "go_race_condition/src/router"
    usecases "go_race_condition/src/usecases"
)

func main() {
    productRepo := repo.NewProductRepository(100000)
    productUsecase := usecases.NewProductUsecase(productRepo)
    handler := handlers.NewHandler(productUsecase)

    r := router.NewRouter(handler)

    fmt.Println("Server is running on port 8080")
    log.Fatal(http.ListenAndServe(":8080", r))
}
