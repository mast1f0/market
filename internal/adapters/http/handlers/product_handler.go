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
	ctx := r.Context()
	products, err := h.products.GetAllProducts(ctx)
	if err != nil {
		helpers.RespondErr(w, err)
		return
	}
	helpers.RespondJSON(w, http.StatusOK, products)
}

func (h *ProductHandler) AddProduct(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userID, ok := ctx.Value("user_id").(int64)
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

	product, err := h.products.CreateProduct(ctx, req.ToDomain(userID))
	if err != nil {
		helpers.RespondErr(w, err)
		return
	}

	helpers.RespondJSON(w, http.StatusCreated, product)
}

func (h *ProductHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userID, ok := ctx.Value("user_id").(int64)
	if !ok {
		helpers.RespondError(w, http.StatusUnauthorized, "cannot get user id")
		return
	}

	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		helpers.RespondError(w, http.StatusBadRequest, "invalid id")
		return
	}

	role, ok := r.Context().Value("role").(string)
	if !ok {
		helpers.RespondError(w, http.StatusUnauthorized, "cannot get user role")
		return
	}

	if err := h.products.DeleteProduct(ctx, id, userID, role); err != nil {
		helpers.RespondErr(w, err)
		return
	}

	helpers.RespondJSON(w, http.StatusNoContent, nil)
}

func (h *ProductHandler) PutProduct(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userID, ok := ctx.Value("user_id").(int64)
	if !ok {
		helpers.RespondError(w, http.StatusUnauthorized, "cannot get user id")
		return
	}

	role, ok := ctx.Value("role").(string)
	if !ok {
		helpers.RespondError(w, http.StatusUnauthorized, "cannot get user role")
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

	updated, err := h.products.UpdateProduct(ctx, req.ToDomain(userID, id), role)
	if err != nil {
		helpers.RespondErr(w, err)
		return
	}

	helpers.RespondJSON(w, http.StatusOK, updated)
}

func (h *ProductHandler) GetProductById(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)

	product, err := h.products.GetProductById(ctx, id)
	if err != nil {
		helpers.RespondErr(w, err)
		return
	}

	helpers.RespondJSON(w, http.StatusOK, product)
}
