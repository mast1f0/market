package handlers

import (
	"encoding/json"
	"market/internal/adapters/http/dto"
	"net/http"

	"market/internal/adapters/http/helpers"
	"market/internal/core/service"
)

type CartHandler struct {
	service *service.CartService
}

func NewCartHandler(service *service.CartService) *CartHandler {
	return &CartHandler{service: service}
}

func (h *CartHandler) GetCart(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value("user_id").(int64)
	if !ok {
		helpers.RespondError(w, http.StatusUnauthorized, "cannot get user id")
		return
	}

	cart, err := h.service.GetCartWithItems(userID)
	if err != nil {
		helpers.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	helpers.RespondJSON(w, http.StatusOK, cart)
}

func (h *CartHandler) AddItem(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value("user_id").(int64)
	if !ok {
		helpers.RespondError(w, http.StatusUnauthorized, "cannot get user id")
		return
	}
	var req dto.AddCartItemRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		helpers.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}

	newItem, err := h.service.AddCartItem(userID, req.ToDomain())
	if err != nil {
		helpers.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}
	helpers.RespondJSON(w, http.StatusCreated, newItem)
}
func (h *CartHandler) RemoveItem(w http.ResponseWriter, r *http.Request) {
	userId, ok := r.Context().Value("user_id").(int64)
	if !ok {
		helpers.RespondError(w, http.StatusUnauthorized, "cannot get user id")
		return
	}
	var item dto.RemoveCartItemRequest
	err := json.NewDecoder(r.Body).Decode(&item)
	if err != nil {
		helpers.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}
	err = h.service.DeleteCartItem(userId, item.ItemID)
	if err != nil {
		helpers.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}
	helpers.RespondJSON(w, http.StatusOK, item)
}

func (h *CartHandler) UpdateItem(w http.ResponseWriter, r *http.Request) {
	var req dto.UpdateCartItemRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		helpers.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}
	newCart, err := h.service.UpdateCartItem(req.ItemID, req.Quantity)
	if err != nil {
		helpers.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}
	helpers.RespondJSON(w, http.StatusOK, newCart)
}
