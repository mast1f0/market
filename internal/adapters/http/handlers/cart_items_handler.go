package handlers

import (
	"encoding/json"
	"market/internal/core/domain"
	"market/internal/core/service"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type CartItemsHandler struct {
	repo *service.CartItemsService
}

func NewCartItemsHandler(repo *service.CartItemsService) *CartItemsHandler {
	return &CartItemsHandler{
		repo: repo,
	}
}

func (h *CartItemsHandler) AddItemCart(w http.ResponseWriter, r *http.Request) {
	var CartItem domain.CartItems
	err := json.NewDecoder(r.Body).Decode(&CartItem)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	cartItemJson, _ := json.Marshal(CartItem)
	w.WriteHeader(http.StatusCreated)
	w.Write(cartItemJson)
}

func (h *CartItemsHandler) DeleteItemCart(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("invalid id"))
		return
	}
	err = h.repo.DeleteCartItem(int64(id))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *CartItemsHandler) GetCartItems(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("invalid id"))
		return
	}
	items, err := h.repo.GetCartItems(int64(id))
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(items)
}

func (h *CartItemsHandler) UpdateCartItem(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("invalid id"))
		return
	}
	var CartItem domain.CartItems
	err = json.NewDecoder(r.Body).Decode(&CartItem)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	CartItem.Id = int64(id)
	_, err = h.repo.UpdateCartItem(&CartItem)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	cartItemJson, _ := json.Marshal(CartItem)
	w.WriteHeader(http.StatusCreated)
	w.Write(cartItemJson)
}
