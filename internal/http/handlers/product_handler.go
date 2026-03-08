package handlers

import (
	"encoding/json"
	"fmt"
	"market/internal/domain"
	"market/internal/service"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type ProductHandler struct {
	service *service.ProductService
}

func NewProductHandler(s *service.ProductService) *ProductHandler {
	return &ProductHandler{service: s}
}

func (h *ProductHandler) GetAllProducts(w http.ResponseWriter, r *http.Request) {
	products := h.service.GetAllProducts()
	for _, product := range products {
		fmt.Fprintln(w, product)
	}
}

func (h *ProductHandler) AddProduct(w http.ResponseWriter, r *http.Request) {
	var Product *domain.Product
	var err error
	err = json.NewDecoder(r.Body).Decode(&Product)
	if err != nil {
		http.Error(w, "Incorrect body", http.StatusBadRequest)
		return
	}
	product, err := h.service.AddToProduct(Product)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("unable to add product"))
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	jsonProduct, _ := json.Marshal(product)
	_, err = w.Write(jsonProduct)
	if err != nil {
		return
	}

}

func (h *ProductHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("invalid id"))
		return
	}
	var product *domain.Product
	product = h.service.GetProductById(id)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("invalid id"))
		return
	}
	productJson, err := json.Marshal(product)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("unable to get product"))
		return
	}
	w.Write(productJson)
	h.service.DeleteProduct(id)

}

func (h *ProductHandler) PutProduct(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("invalid id"))
		return
	}
	var product *domain.Product
	product.ID = id
	err = json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		http.Error(w, "Incorrect body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()
	productJson, err := json.Marshal(product)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("unable to get product"))
		return
	}
	w.Write(productJson)
	h.service.UpdateProduct(product)
}
