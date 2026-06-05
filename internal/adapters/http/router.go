package http

import (
	"market/internal/adapters/http/handlers"
	"market/internal/adapters/http/middleware"
	jwtutil "market/internal/adapters/jwt"
	"market/internal/engine/logger"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

type AllHandler struct {
	ProductHandler  *handlers.ProductHandler
	CategoryHandler *handlers.CategoryHandler
	CartHandler     *handlers.CartHandler
	OrderHandler    *handlers.OrderHandler
}

func SetupRoutes(handlers *AllHandler, jwt *jwtutil.Manager, logger *logger.Logger) *chi.Mux {
	r := chi.NewRouter()

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: false,
	}))
	r.Use(middleware.LoggingMiddleware(logger))
	//не требует авторизации
	r.Group(func(r chi.Router) {
		r.Get("/products", handlers.ProductHandler.GetAllProducts)
		r.Get("/products/{id}", handlers.ProductHandler.GetProductById)
		r.Get("/products/search", handlers.ProductHandler.GetProductsByName)

		r.Get("/categories/{id}", handlers.CategoryHandler.ListProductsByCategoryID)
		r.Get("/categories", handlers.CategoryHandler.ListCategories)
	})

	//авторизованные типочки
	r.Group(func(r chi.Router) {
		r.Use(middleware.AuthMiddleware(jwt))

		r.Group(func(r chi.Router) {
			r.Use(middleware.RoleMiddleware("seller", "admin"))

			r.Post("/products", handlers.ProductHandler.AddProduct)
			r.Delete("/products/{id}", handlers.ProductHandler.DeleteProduct)
			r.Put("/products/{id}", handlers.ProductHandler.PutProduct)

			r.Post("/categories", handlers.CategoryHandler.AddCategory)
			r.Delete("/categories/{id}", handlers.CategoryHandler.DeleteCategory)
			r.Put("/categories/{id}", handlers.CategoryHandler.UpdateCategory)

			r.Put("/orders/{id}", handlers.OrderHandler.UpdateOrder)
		})

		r.Group(func(r chi.Router) {
			r.Use(middleware.RoleMiddleware("buyer", "seller", "admin"))
			r.Get("/cart", handlers.CartHandler.GetCart)
			r.Post("/cart/items", handlers.CartHandler.AddItem)
			r.Delete("/cart/items", handlers.CartHandler.RemoveItem)
			r.Put("/cart/items", handlers.CartHandler.UpdateItem)

			r.Get("/orders", handlers.OrderHandler.GetOrderByUser)
			r.Get("/orders/{id}", handlers.OrderHandler.GetOrderById)
			r.Post("/orders", handlers.OrderHandler.CreateOrder)
		})
	})

	return r
}
