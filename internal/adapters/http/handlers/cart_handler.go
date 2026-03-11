package handlers

import (
	"encoding/json"
	"market/internal/core/domain"
	"market/internal/core/service"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type CartHandler struct {
	repo *service.CartService
}

func NewCartHandler(repo *service.CartService) *CartHandler {
	return &CartHandler{repo: repo}
}

func (h *CartHandler) CreateCart(w http.ResponseWriter, r *http.Request) {
	var Cart domain.Cart
	err := json.NewDecoder(r.Body).Decode(&Cart)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var newCart *domain.Cart
	newCart, err = h.repo.CreateCart(&Cart)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	newCartJson, _ := json.Marshal(newCart)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(newCartJson)
}

func (h *CartHandler) GetCart(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	cart, err := h.repo.GetCart(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	cartJson, _ := json.Marshal(cart)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(cartJson)
}

func (h *CartHandler) UpdateCart(w http.ResponseWriter, r *http.Request) {
	var Cart domain.Cart
	err := json.NewDecoder(r.Body).Decode(&Cart)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	Cart.Id = id
	var newCart *domain.Cart
	newCart, err = h.repo.UpdateCart(&Cart)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	newCartJson, _ := json.Marshal(newCart)
	w.Write(newCartJson)
}

func (h *CartHandler) DeleteCart(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = h.repo.DeleteCart(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}
