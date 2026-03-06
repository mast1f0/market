package http

import (
	"encoding/json"
	"fmt"
	"market/internal/domain"
	"market/internal/ports"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type Handlers struct {
	Router  *chi.Mux
	Storage ports.StorageRepository
}

func NewHandlers(stor ports.StorageRepository) *Handlers {
	r := chi.NewRouter()
	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "OK")
	})
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		products := stor.GetAllProducts()
		for _, product := range products {
			fmt.Fprintln(w, product)
		}
	})
	r.HandleFunc("POST /products/{id}", func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		var Product domain.Product
		Product.ID = id
		_ = json.NewDecoder(r.Body).Decode(&Product)
		product, err := stor.AddProduct(Product)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusBadRequest)
		}
		w.WriteHeader(http.StatusCreated)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, product)

	})
	return &Handlers{
		Router:  r,
		Storage: stor,
	}
}
