package main

import (
	"market/internal/adapters/storage/memory"
	"market/internal/http"
	"market/internal/http/handlers"
	"market/internal/service"
)

func main() {
	repo := memory.NewMemory()

	productService := service.NewProductService(repo)
	productHandler := handlers.NewProductHandler(productService)
	router := http.SetupRoutes(productHandler)
	srv := http.NewServer(router)
	srv.Start()
}
