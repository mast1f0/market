package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"market/internal/adapters/http/dto"
	"market/internal/adapters/http/helpers"
	"market/internal/core/domain"
	"market/internal/core/service"

	"github.com/go-chi/chi/v5"
)

type CategoryHandler struct {
	service *service.CategoryService
}

func NewCategoryHandler(service *service.CategoryService) *CategoryHandler {
	return &CategoryHandler{service: service}
}

func (h *CategoryHandler) AddCategory(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateCategoryRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		helpers.RespondError(w, http.StatusBadRequest, "invalid json body")
		return
	}
	if err := req.Validate(); err != nil {
		helpers.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}

	created, err := h.service.CreateCategory(strings.TrimSpace(req.Name))
	if err != nil {
		helpers.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	helpers.RespondJSON(w, http.StatusCreated, created)
}

func (h *CategoryHandler) GetCategory(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	category, err := h.service.GetCategory(id)
	if err != nil {
		helpers.RespondError(w, http.StatusNotFound, err.Error())
		return
	}

	helpers.RespondJSON(w, http.StatusOK, category)
}

func (h *CategoryHandler) DeleteCategory(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err = h.service.DeleteCategory(id); err != nil {
		//со статусами надо порешать
		helpers.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	helpers.RespondJSON(w, http.StatusNoContent, nil)
}

func (h *CategoryHandler) UpdateCategory(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	var req dto.UpdateCategoryRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		helpers.RespondError(w, http.StatusBadRequest, "invalid json body")
		return
	}
	if err := req.Validate(); err != nil {
		helpers.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}

	if _, err := h.service.GetCategory(id); err != nil {
		helpers.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	category := &domain.Category{
		ID:   id,
		Name: strings.TrimSpace(req.Name),
	}

	updated, err := h.service.UpdateCategory(category)
	if err != nil {
		helpers.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}

	helpers.RespondJSON(w, http.StatusOK, updated)
}

func (h *CategoryHandler) ListCategories(w http.ResponseWriter, r *http.Request) {
	categories, err := h.service.GetCategories()
	if err != nil {
		helpers.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	helpers.RespondJSON(w, http.StatusOK, categories)
}

func (h *CategoryHandler) ListProductsByCategoryID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 32)
	products, err := h.service.GetCategoriesByCategoryID(id)
	if err != nil {
		helpers.RespondError(w, http.StatusNotFound, err.Error())
		return
	}
	helpers.RespondJSON(w, http.StatusOK, products)
}
