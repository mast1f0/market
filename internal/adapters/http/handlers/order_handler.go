package handlers

import (
	"encoding/json"
	"market/internal/adapters/http/dto"
	"market/internal/adapters/http/helpers"
	"market/internal/core/service"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type OrderHandler struct {
	service *service.OrderService
}

func NewOrderHandler(service *service.OrderService) *OrderHandler {
	return &OrderHandler{
		service: service,
	}
}

func (h *OrderHandler) GetOrderByUser(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value("user_id").(int64)
	if !ok {
		helpers.RespondError(w, http.StatusUnauthorized, "cannot get user id")
		return
	}

	orders, err := h.service.GetByUserId(userID)
	if err != nil {
		helpers.RespondError(w, http.StatusBadRequest, "cannot get orders")
		return
	}
	helpers.RespondJSON(w, http.StatusOK, orders)
}

func (h *OrderHandler) GetOrderById(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value("user_id").(int64)
	if !ok {
		helpers.RespondError(w, http.StatusUnauthorized, "cannot get user id")
		return
	}
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		helpers.RespondError(w, http.StatusBadRequest, "cannot get order")
		return
	}
	order, err := h.service.GetOrderById(id)
	if order.UserId != userID {
		helpers.RespondError(w, http.StatusUnauthorized, "cannot get order")
		return
	}

	if err != nil {
		helpers.RespondError(w, http.StatusInternalServerError, "cannot get order")
		return
	}
	helpers.RespondJSON(w, http.StatusOK, order)

}

func (h *OrderHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value("user_id").(int64)
	if !ok {
		helpers.RespondError(w, http.StatusUnauthorized, "cannot get user id")
		return
	}
	order, err := h.service.CreateFromCart(userID)
	if err != nil {
		helpers.RespondError(w, http.StatusBadRequest, "cannot create order")
		return
	}
	helpers.RespondJSON(w, http.StatusCreated, order)
}

func (h *OrderHandler) UpdateOrder(w http.ResponseWriter, r *http.Request) {
	_, ok := r.Context().Value("user_id").(int64)
	if !ok {
		helpers.RespondError(w, http.StatusUnauthorized, "cannot get user id")
		return
	}
	strId := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(strId, 10, 64)
	if err != nil {
		helpers.RespondError(w, http.StatusBadRequest, "cannot get order")
		return
	}
	var req dto.UpdateOrderStatusRequest
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		helpers.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}
	err = h.service.UpdateStatus(id, req.Status)
	if err != nil {
		helpers.RespondError(w, http.StatusInternalServerError, "cannot update order")
		return
	}
	helpers.RespondJSON(w, http.StatusOK, map[string]string{
		"status": "updated",
	})
}
