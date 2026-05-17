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

type ProductHandler struct {
	products *service.ProductService
}

func NewProductHandler(products *service.ProductService) *ProductHandler {
	return &ProductHandler{products: products}
}

func (h *ProductHandler) GetAllProducts(w http.ResponseWriter, r *http.Request) {
	products, err := h.products.GetAllProducts()
	if err != nil {
		helpers.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	helpers.RespondJSON(w, http.StatusOK, products)
}

func (h *ProductHandler) AddProduct(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value("user_id").(int64)
	if !ok {
		helpers.RespondError(w, http.StatusUnauthorized, "cannot get user id")
		return
	}

	var req dto.CreateProductRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		helpers.RespondError(w, http.StatusBadRequest, "invalid json body")
		return
	}
	if err := req.Validate(); err != nil {
		helpers.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}

	product, err := h.products.CreateProduct(req.ToDomain(userID))
	if err != nil {
		helpers.RespondError(w, http.StatusBadRequest, "unable to add product")
		return
	}

	helpers.RespondJSON(w, http.StatusCreated, product)
}

func (h *ProductHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value("user_id").(int64)
	if !ok {
		helpers.RespondError(w, http.StatusUnauthorized, "cannot get user id")
		return
	}

	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		helpers.RespondError(w, http.StatusBadRequest, "invalid id")
		return
	}

	if err := h.products.DeleteProduct(id, userID); err != nil {
		helpers.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	helpers.RespondJSON(w, http.StatusNoContent, nil)
}

func (h *ProductHandler) PutProduct(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value("user_id").(int64)
	if !ok {
		helpers.RespondError(w, http.StatusUnauthorized, "cannot get user id")
		return
	}

	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		helpers.RespondError(w, http.StatusBadRequest, "invalid id")
		return
	}
	var req dto.UpdateProductRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		helpers.RespondError(w, http.StatusBadRequest, "invalid json body")
		return
	}

	updated, err := h.products.UpdateProduct(req.ToDomain(userID, id))
	if err != nil {
		helpers.RespondError(w, http.StatusBadRequest, "unable to update product")
		return
	}

	helpers.RespondJSON(w, http.StatusOK, updated)
}

func (h *ProductHandler) GetProductById(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)

	product, err := h.products.GetProductById(id)
	if err != nil {
		helpers.RespondError(w, http.StatusNotFound, err.Error())
		return
	}

	helpers.RespondJSON(w, http.StatusOK, product)
}
