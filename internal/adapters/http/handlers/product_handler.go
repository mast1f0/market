package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"market/internal/adapters/http/dto"
	"market/internal/adapters/http/helpers"
	"market/internal/core/service"

	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
)

type ProductHandler struct {
	products *service.ProductService
}

func NewProductHandler(products *service.ProductService, categories *service.CategoryService) *ProductHandler {
	return &ProductHandler{products: products}
}

func (h *ProductHandler) GetAllProducts(w http.ResponseWriter, r *http.Request) {
	products := h.products.GetAllProducts()
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

	product, err := h.products.AddToProduct(req.ToDomain(userID))
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
	if err != nil || id < 1 {
		helpers.RespondError(w, http.StatusBadRequest, "invalid id")
		return
	}

	product, err := h.products.GetProductById(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) || helpers.HTTPStatusForDB(err) == http.StatusNotFound {
			helpers.RespondError(w, http.StatusNotFound, "product not found")
			return
		}
		helpers.RespondError(w, http.StatusInternalServerError, "failed to load product")
		return
	}

	if product.OwnerID != userID {
		helpers.RespondError(w, http.StatusForbidden, "you are not the owner of this product")
		return
	}

	if err := h.products.DeleteProduct(id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helpers.RespondError(w, http.StatusNotFound, "product not found")
			return
		}
		helpers.RespondError(w, http.StatusInternalServerError, "failed to delete product")
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
	updated, err := h.products.UpdateProduct(id, userID)
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
		if errors.Is(err, gorm.ErrRecordNotFound) || helpers.HTTPStatusForDB(err) == http.StatusNotFound {
			helpers.RespondError(w, http.StatusNotFound, "product not found")
			return
		}
		helpers.RespondError(w, http.StatusInternalServerError, "failed to load product")
		return
	}

	helpers.RespondJSON(w, http.StatusOK, product)
}
