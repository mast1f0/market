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

	db := orm.NewConnection()
	productRepository := orm.NewProductRepository(db)
	categoryRepository := orm.NewCategoryRepository(db)
	cartRepository := orm.NewCartRepository(db)
	orderRepository := orm.NewOrderRepository(db)

	if *seedFlag {
		seed2.SeedCategories(categoryRepository)
		seed2.SeedProducts(productRepository, categoryRepository)
	}
	productService := service.NewProductService(productRepository)
	categoryService := service.NewCategoryService(categoryRepository)
	cartService := service.NewCartService(cartRepository)
	orderService := service.NewOrderService(orderRepository, cartRepository)

	productHandler := handlers.NewProductHandler(productService, categoryService)
	categoryHandler := handlers.NewCategoryHandler(categoryService)
	cartHandler := handlers.NewCartHandler(cartService)
	orderHanler := handlers.NewOrderHandler(orderService)

	jwt := jwtutil.Manager{Secret: []byte("superSecret")}
	router := http2.SetupRoutes(productHandler, categoryHandler, cartHandler, orderHanler, &jwt)
	srv := http2.NewServer(router)
	srv.Start()
}
