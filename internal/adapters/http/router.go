package http

import (
	"market/internal/adapters/http/handlers"
	"market/internal/adapters/http/middlware"
	jwtutil "market/internal/adapters/jwt"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func SetupRoutes(productHandler *handlers.ProductHandler, categoryHandler *handlers.CategoryHandler, cartHandler *handlers.CartHandler, CartItemsHandler *handlers.CartItemsHandler, jwt *jwtutil.Manager) *chi.Mux { //, CategoryHandler *handlers.CategoryHandler
	r := chi.NewRouter()

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: false,
	}))
	r.Use(middlware.AuthMiddleware(jwt))

	//не требует авторизации
	r.Group(func(r chi.Router) {
		r.Get("/products", productHandler.GetAllProducts)
		r.Get("/products/{id}", productHandler.GetProductById)

		r.Get("/categories/{id}", categoryHandler.GetCategory)
		r.Get("/categories", categoryHandler.ListCategories)
	})

	//авторизованные типочки
	r.Group(func(r chi.Router) {
		r.Use(middlware.AuthMiddleware(jwt))

		r.Get("/carts/{id}", cartHandler.GetCart)
		r.Post("/carts", cartHandler.CreateCart)
		r.Put("/carts/{id}", cartHandler.UpdateCart)
		r.Delete("/carts/{id}", cartHandler.DeleteCart)

		r.Get("/cartsItems/{id}", CartItemsHandler.GetCartItems)
		r.Post("/cartsItems/{id}", CartItemsHandler.AddItemCart)
		r.Delete("/cartsItems/{id}", CartItemsHandler.DeleteItemCart)
		r.Put("/cartsItems/{id}", CartItemsHandler.UpdateCartItem)

		r.Group(func(r chi.Router) {
			r.Use(middlware.RoleMiddleware("buyer", "seller", "admin"))

			r.Post("/products", productHandler.AddProduct)
			r.Delete("/products/{id}", productHandler.DeleteProduct)
			r.Put("/products/{id}", productHandler.PutProduct)

			r.Post("/categories", categoryHandler.AddCategory)
			r.Delete("/categories/{id}", categoryHandler.DeleteCategory)
			r.Put("/categories/{id}", categoryHandler.UpdateCategory)
		})

	})

	return r
}
