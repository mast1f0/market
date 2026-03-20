package main

import (
	http2 "market/internal/adapters/http"
	"market/internal/adapters/http/handlers"
	jwtutil "market/internal/adapters/jwt"
	"market/internal/adapters/storage/orm"
	"market/internal/core/service"
	"market/internal/logger"

	log "github.com/sirupsen/logrus"
)

func main() {
	logger.Init()
	repo := orm.NewStorage()

	log.Debug("init storage")
	productService := service.NewProductService(repo)
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
