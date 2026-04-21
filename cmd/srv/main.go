package main

import (
	"flag"
	http2 "market/internal/adapters/http"
	"market/internal/adapters/http/handlers"
	jwtutil "market/internal/adapters/jwt"
	"market/internal/adapters/storage/orm"
	"market/internal/core/service"
	"market/internal/engine/logger"
	seed2 "market/internal/engine/seed"
)

func main() {
	logger.Init()
	seedFlag := flag.Bool("seed", false, "run seed")
	flag.Parse()

	repo := orm.NewStorage()
	if *seedFlag {
		seed2.SeedCategories(repo)
		seed2.SeedProducts(repo)
	}
	productService := service.NewProductService(repo)
	categoryService := service.NewCategoryService(repo)
	cartService := service.NewCartService(repo)
	cartItemsService := service.NewCartItemsService(repo)

	productHandler := handlers.NewProductHandler(productService, categoryService)
	categoryHandler := handlers.NewCategoryHandler(categoryService)
	cartHandler := handlers.NewCartHandler(cartService)
	cartItemsHandler := handlers.NewCartItemsHandler(cartItemsService, cartService)

	jwt := jwtutil.Manager{Secret: []byte("superSecret")}
	router := http2.SetupRoutes(productHandler, categoryHandler, cartHandler, cartItemsHandler, &jwt)
	srv := http2.NewServer(router)
	srv.Start()
}
