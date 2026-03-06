package http

import (
	"market/internal/http/handlers"

	"github.com/go-chi/chi/v5"
)

type Handlers struct {
	Storage handlers.ProductHandler
}

func SetupRoutes(productHandler *handlers.ProductHandler) *chi.Mux {
	r := chi.NewRouter()
	r.Get("/products/", productHandler.GetAllProducts)
	r.Post("/products/", productHandler.AddProduct)
	r.Delete("/products/{id}", productHandler.DeleteProduct)
	r.Put("/products/{id}", productHandler.PutProduct)
	return r
}
