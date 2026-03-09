package main

import (
	http2 "market/internal/adapters/http"
	"market/internal/adapters/http/handlers"
	"market/internal/adapters/storage/orm"
	"market/internal/core/service"
)

func main() {
	repo := orm.NewStorage()

	productService := service.NewProductService(repo)
	//categoryService := service.NewCategoryService(repo)
	productHandler := handlers.NewProductHandler(productService)
	//categoryHandler := handlers.NewCategoryHandler(categoryService)
	router := http2.SetupRoutes(productHandler)
	srv := http2.NewServer(router)
	srv.Start()
}
