package main

import (
	"market/internal/adapters/storage/orm"
	"market/internal/http"
	"market/internal/http/handlers"
	"market/internal/service"
)

func main() {
	repo := orm.NewStorage()

	productService := service.NewProductService(repo)
	//categoryService := service.NewCategoryService(repo)
	productHandler := handlers.NewProductHandler(productService)
	//categoryHandler := handlers.NewCategoryHandler(categoryService)
	router := http.SetupRoutes(productHandler)
	srv := http.NewServer(router)
	srv.Start()
}
