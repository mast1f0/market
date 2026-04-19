package handlers

import (
	"encoding/json"
	"net/http"

	"market/internal/adapters/http/dto"
	"market/internal/adapters/http/helpers"
	"market/internal/core/domain"
	"market/internal/core/service"
)

type CartHandler struct {
	repo *service.CartService
}

func NewCartHandler(repo *service.CartService) *CartHandler {
	return &CartHandler{repo: repo}
}

func (h *CartHandler) CreateCart(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value("user_id").(int64)
	if !ok {
		helpers.RespondError(w, http.StatusUnauthorized, "cannot get user id")
		return
	}

	newCart, err := h.repo.CreateCart(&domain.Cart{UserId: userID})
	if err != nil {
		helpers.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}

	helpers.RespondJSON(w, http.StatusCreated, newCart)
}

func (h *CartHandler) GetCart(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value("user_id").(int64)
	if !ok {
		helpers.RespondError(w, http.StatusUnauthorized, "cannot get user id")
		return
	}

	cart, err := h.repo.GetCart(userID)
	if err != nil {
		if helpers.HTTPStatusForDB(err) == http.StatusNotFound {
			helpers.RespondError(w, http.StatusNotFound, "cart not found")
			return
		}
		helpers.RespondError(w, http.StatusInternalServerError, "failed to load cart")
		return
	}

	helpers.RespondJSON(w, http.StatusOK, cart)
}

func (h *CartHandler) UpdateCart(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value("user_id").(int64)
	if !ok {
		helpers.RespondError(w, http.StatusUnauthorized, "cannot get user id")
		return
	}

	cart, err := h.repo.GetCart(userID)
	if err != nil {
		if helpers.HTTPStatusForDB(err) == http.StatusNotFound {
			helpers.RespondError(w, http.StatusNotFound, "cart not found")
			return
		}
		helpers.RespondError(w, http.StatusInternalServerError, "failed to load cart")
		return
	}

	var req dto.UpdateCartRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		helpers.RespondError(w, http.StatusBadRequest, "invalid json body")
		return
	}

	updated, err := h.repo.UpdateCart(cart)
	if err != nil {
		helpers.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}

	helpers.RespondJSON(w, http.StatusOK, updated)
}

func (h *CartHandler) DeleteCart(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value("user_id").(int64)
	if !ok {
		helpers.RespondError(w, http.StatusUnauthorized, "cannot get user id")
		return
	}

	if err := h.repo.DeleteCart(userID); err != nil {
		if helpers.HTTPStatusForDB(err) == http.StatusNotFound {
			helpers.RespondError(w, http.StatusNotFound, "cart not found")
			return
		}
		helpers.RespondError(w, http.StatusInternalServerError, "failed to delete cart")
		return
	}

	helpers.RespondJSON(w, http.StatusNoContent, nil)
}
