package main

import (
	"flag"
	http2 "market/internal/adapters/http"
	"market/internal/adapters/http/handlers"
	jwtutil "market/internal/adapters/jwt"
	"market/internal/adapters/storage/orm"
	s3storage "market/internal/adapters/storage/s3"
	"market/internal/core/service"
	"market/internal/logger"
	"market/internal/seed"
)

func main() {
	logger.Init()
	seedFlag := flag.Bool("seed", false, "run seed")
	flag.Parse()

	repo := orm.NewStorage()
	if *seedFlag {
		seed.SeedCategories(repo)
		seed.SeedProducts(repo)
	}
	s3 := s3storage.NewS3Client("BUCKET")

	s3Storage := service.NewS3Service(s3)
	productService := service.NewProductService(repo, s3Storage)
	categoryService := service.NewCategoryService(repo)
	cartService := service.NewCartService(repo)
	cartItemsService := service.NewCartItemsService(repo)

	productHandler := handlers.NewProductHandler(productService)
	categoryHandler := handlers.NewCategoryHandler(categoryService)
	cartHandler := handlers.NewCartHandler(cartService)
	cartItemsHandler := handlers.NewCartItemsHandler(cartItemsService)

	jwt := jwtutil.Manager{Secret: []byte("superSecret")}
	router := http2.SetupRoutes(productHandler, categoryHandler, cartHandler, cartItemsHandler, &jwt)
	srv := http2.NewServer(router)
	srv.Start()
}
