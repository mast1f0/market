package http

import (
	"market/internal/adapters/http/handlers"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func SetupRoutes(productHandler *handlers.ProductHandler, categoryHandler *handlers.CategoryHandler, cartHandler *handlers.CartHandler, CartItemsHandler *handlers.CartItemsHandler) *chi.Mux { //, CategoryHandler *handlers.CategoryHandler
	r := chi.NewRouter()

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: false,
	}))
	r.Get("/products", productHandler.GetAllProducts)
	r.Post("/products", productHandler.AddProduct)
	r.Delete("/products/{id}", productHandler.DeleteProduct)
	r.Put("/products/{id}", productHandler.PutProduct)
	
	r.Get("/categories/{id}", categoryHandler.GetCategory)
	r.Post("/categories", categoryHandler.AddCategory)
	r.Delete("/categories/{id}", categoryHandler.DeleteCategory)
	r.Put("/categories/{id}", categoryHandler.UpdateCategory)

	r.Get("/carts/{id}", cartHandler.GetCart)
	r.Post("/carts", cartHandler.CreateCart)
	r.Put("/carts/{id}", cartHandler.UpdateCart)
	r.Delete("/carts/{id}", cartHandler.DeleteCart)

	r.Get("/carts/{id}", CartItemsHandler.GetCartItems)
	r.Post("/carts/{id}", CartItemsHandler.AddItemCart)
	r.Delete("/carts/{id}", CartItemsHandler.DeleteItemCart)
	r.Put("/carts/{id}", CartItemsHandler.UpdateCartItem)
	return r
}
