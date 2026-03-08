package http

import (
	"market/internal/http/handlers"

	"github.com/go-chi/chi/v5"
)

func SetupRoutes(productHandler *handlers.ProductHandler) *chi.Mux { //, CategoryHandler *handlers.CategoryHandler
	r := chi.NewRouter()
	r.Get("/products/", productHandler.GetAllProducts)
	r.Post("/products/", productHandler.AddProduct)
	r.Delete("/products/{id}", productHandler.DeleteProduct)
	r.Put("/products/{id}", productHandler.PutProduct)

	//r.Get("/categories/", CategoryHandler.GetAllCategories)
	//r.Post("/categories/", CategoryHandler.AddCategory)
	//r.Get("/categories/{id}", CategoryHandler.GetCategoryById)
	//r.Delete("/categories/{id}", CategoryHandler.DeleteCategoryById)
	//r.Put("/categories/{id}", CategoryHandler.Update)
	return r
}
