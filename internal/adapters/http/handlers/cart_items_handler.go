package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"market/internal/adapters/http/dto"
	"market/internal/adapters/http/helpers"
	"market/internal/core/service"

	"github.com/go-chi/chi/v5"
)

type CartItemsHandler struct {
	items *service.CartItemsService
	carts *service.CartService
}

func NewCartItemsHandler(items *service.CartItemsService, carts *service.CartService) *CartItemsHandler {
	return &CartItemsHandler{items: items, carts: carts}
}

func (h *CartItemsHandler) AddItemCart(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value("user_id").(int64)
	if !ok {
		helpers.RespondError(w, http.StatusUnauthorized, "cannot get user id")
		return
	}

	cartID, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil || cartID < 1 {
		helpers.RespondError(w, http.StatusBadRequest, "invalid cart id")
		return
	}

	cart, err := h.carts.GetCartByID(cartID)
	if err != nil {
		if helpers.HTTPStatusForDB(err) == http.StatusNotFound {
			helpers.RespondError(w, http.StatusNotFound, "cart not found")
			return
		}
		helpers.RespondError(w, http.StatusInternalServerError, "failed to load cart")
		return
	}
	if cart.UserId != userID {
		helpers.RespondError(w, http.StatusForbidden, "cart does not belong to current user")
		return
	}

	var req dto.AddCartItemRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		helpers.RespondError(w, http.StatusBadRequest, "invalid json body")
		return
	}
	if err := req.Validate(); err != nil {
		helpers.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}

	created, err := h.items.AddCartItem(req.ToDomain(cartID))
	if err != nil {
		helpers.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}

	helpers.RespondJSON(w, http.StatusCreated, created)
}

func (h *CartItemsHandler) DeleteItemCart(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value("user_id").(int64)
	if !ok {
		helpers.RespondError(w, http.StatusUnauthorized, "cannot get user id")
		return
	}

	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil || id < 1 {
		helpers.RespondError(w, http.StatusBadRequest, "invalid id")
		return
	}

	item, err := h.items.GetCartItems(id)
	if err != nil {
		if helpers.HTTPStatusForDB(err) == http.StatusNotFound {
			helpers.RespondError(w, http.StatusNotFound, "cart item not found")
			return
		}
		helpers.RespondError(w, http.StatusInternalServerError, "failed to load cart item")
		return
	}

	cart, err := h.carts.GetCartByID(item.CartId)
	if err != nil || cart.UserId != userID {
		helpers.RespondError(w, http.StatusForbidden, "cannot delete this cart item")
		return
	}

	if err := h.items.DeleteCartItem(id); err != nil {
		if helpers.HTTPStatusForDB(err) == http.StatusNotFound {
			helpers.RespondError(w, http.StatusNotFound, "cart item not found")
			return
		}
		helpers.RespondError(w, http.StatusInternalServerError, "failed to delete cart item")
		return
	}

	helpers.RespondJSON(w, http.StatusNoContent, nil)
}

func (h *CartItemsHandler) GetCartItems(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value("user_id").(int64)
	if !ok {
		helpers.RespondError(w, http.StatusUnauthorized, "cannot get user id")
		return
	}

	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil || id < 1 {
		helpers.RespondError(w, http.StatusBadRequest, "invalid id")
		return
	}

	items, err := h.items.GetCartItems(id)
	if err != nil {
		if helpers.HTTPStatusForDB(err) == http.StatusNotFound {
			helpers.RespondError(w, http.StatusNotFound, "cart item not found")
			return
		}
		helpers.RespondError(w, http.StatusInternalServerError, "failed to load cart item")
		return
	}

	cart, err := h.carts.GetCartByID(items.CartId)
	if err != nil || cart.UserId != userID {
		helpers.RespondError(w, http.StatusForbidden, "cannot access this cart item")
		return
	}

	helpers.RespondJSON(w, http.StatusOK, items)
}

func (h *CartItemsHandler) UpdateCartItem(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value("user_id").(int64)
	if !ok {
		helpers.RespondError(w, http.StatusUnauthorized, "cannot get user id")
		return
	}

	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil || id < 1 {
		helpers.RespondError(w, http.StatusBadRequest, "invalid id")
		return
	}

	existing, err := h.items.GetCartItems(id)
	if err != nil {
		if helpers.HTTPStatusForDB(err) == http.StatusNotFound {
			helpers.RespondError(w, http.StatusNotFound, "cart item not found")
			return
		}
		helpers.RespondError(w, http.StatusInternalServerError, "failed to load cart item")
		return
	}

	cart, err := h.carts.GetCartByID(existing.CartId)
	if err != nil || cart.UserId != userID {
		helpers.RespondError(w, http.StatusForbidden, "cannot update this cart item")
		return
	}

	var req dto.UpdateCartItemRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		helpers.RespondError(w, http.StatusBadRequest, "invalid json body")
		return
	}
	if err := req.Validate(); err != nil {
		helpers.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}

	item := *existing
	req.ApplyTo(&item)
	item.Id = id

	updated, err := h.items.UpdateCartItem(&item)
	if err != nil {
		helpers.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}

	helpers.RespondJSON(w, http.StatusOK, updated)
}
