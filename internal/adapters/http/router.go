package http

import (
	"market/internal/adapters/http/handlers"

	"github.com/go-chi/chi/v5"
)

func SetupRoutes(productHandler *handlers.ProductHandler, categoryHandler *handlers.CategoryHandler) *chi.Mux { //, CategoryHandler *handlers.CategoryHandler
	r := chi.NewRouter()
	r.Get("/products/", productHandler.GetAllProducts)
	r.Post("/products/", productHandler.AddProduct)
	r.Delete("/products/{id}", productHandler.DeleteProduct)
	r.Put("/products/{id}", productHandler.PutProduct)

	r.Get("/categories/{id}", categoryHandler.GetCategory)
	r.Post("/categories", categoryHandler.AddCategory)
	r.Delete("/categories/{id}", categoryHandler.DeleteCategory)
	r.Put("/categories/{id}", categoryHandler.UpdateCategory)
	return r
}
