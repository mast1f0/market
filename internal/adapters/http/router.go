package http

import (
	"market/internal/adapters/http/handlers"
	"market/internal/adapters/http/middleware"
	jwtutil "market/internal/adapters/jwt"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func SetupRoutes(productHandler *handlers.ProductHandler, categoryHandler *handlers.CategoryHandler, cartHandler *handlers.CartHandler, orderHandler *handlers.OrderHandler, jwt *jwtutil.Manager) *chi.Mux { //, CategoryHandler *handlers.CategoryHandler
	r := chi.NewRouter()

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: false,
	}))

	//не требует авторизации
	r.Group(func(r chi.Router) {
		r.Get("/products", productHandler.GetAllProducts)
		r.Get("/products/{id}", productHandler.GetProductById)

		r.Get("/categories/{id}", categoryHandler.ListProductsByCategoryID)
		r.Get("/categories", categoryHandler.ListCategories)
	})

	//авторизованные типочки
	r.Group(func(r chi.Router) {
		r.Use(middleware.AuthMiddleware(jwt))

		r.Group(func(r chi.Router) {
			r.Use(middleware.RoleMiddleware("seller", "admin"))

			r.Post("/products", productHandler.AddProduct)
			r.Delete("/products/{id}", productHandler.DeleteProduct)
			r.Put("/products/{id}", productHandler.PutProduct)

			r.Post("/categories", categoryHandler.AddCategory)
			r.Delete("/categories/{id}", categoryHandler.DeleteCategory)
			r.Put("/categories/{id}", categoryHandler.UpdateCategory)

			r.Put("/orders/{id}", orderHandler.UpdateOrder)
		})

		r.Group(func(r chi.Router) {
			r.Use(middleware.RoleMiddleware("buyer", "seller", "admin"))
			r.Get("/cart", cartHandler.GetCart)
			r.Post("/cart/items", cartHandler.AddItem)
			r.Delete("/cart/items", cartHandler.RemoveItem)
			r.Put("/cart/items", cartHandler.UpdateItem)

			r.Get("/orders", orderHandler.GetOrderByUser)
			r.Get("/orders/{id}", orderHandler.GetOrderById)
			r.Post("/orders", orderHandler.CreateOrder)
		})
	})

	return r
}
