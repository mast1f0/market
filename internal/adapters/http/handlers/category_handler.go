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

	created, err := h.service.CreateCategory(req.ToDomain())
	if err != nil {
		if helpers.IsDuplicateKey(err) {
			helpers.RespondError(w, http.StatusConflict, "category with this name already exists")
			return
		}
		helpers.RespondError(w, http.StatusInternalServerError, "failed to create category")
		return
	}

	helpers.RespondJSON(w, http.StatusCreated, created)
}

func (h *CategoryHandler) GetCategory(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil || id < 1 {
		helpers.RespondError(w, http.StatusBadRequest, "invalid category id")
		return
	}

	category, err := h.service.GetCategory(id)
	if err != nil {
		st := helpers.HTTPStatusForDB(err)
		if st == http.StatusNotFound {
			helpers.RespondError(w, http.StatusNotFound, "category not found")
			return
		}
		helpers.RespondError(w, st, "failed to load category")
		return
	}

	helpers.RespondJSON(w, http.StatusOK, category)
}

func (h *CategoryHandler) DeleteCategory(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil || id < 1 {
		helpers.RespondError(w, http.StatusBadRequest, "invalid category id")
		return
	}

	if err := h.service.DeleteCategory(id); err != nil {
		st := helpers.HTTPStatusForDB(err)
		if st == http.StatusNotFound {
			helpers.RespondError(w, http.StatusNotFound, "category not found")
			return
		}
		helpers.RespondError(w, http.StatusInternalServerError, "failed to delete category")
		return
	}

	helpers.RespondJSON(w, http.StatusNoContent, nil)
}

func (h *CategoryHandler) UpdateCategory(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil || id < 1 {
		helpers.RespondError(w, http.StatusBadRequest, "invalid category id")
		return
	}

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
		if helpers.HTTPStatusForDB(err) == http.StatusNotFound {
			helpers.RespondError(w, http.StatusNotFound, "category not found")
			return
		}
		helpers.RespondError(w, http.StatusInternalServerError, "failed to load category")
		return
	}

	category := &domain.Category{
		ID:   id,
		Name: strings.TrimSpace(req.Name),
	}

	updated, err := h.service.UpdateCategory(category)
	if err != nil {
		if helpers.IsDuplicateKey(err) {
			helpers.RespondError(w, http.StatusConflict, "category with this name already exists")
			return
		}
		st := helpers.HTTPStatusForDB(err)
		if st == http.StatusNotFound {
			helpers.RespondError(w, http.StatusNotFound, "category not found")
			return
		}
		helpers.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}

	helpers.RespondJSON(w, http.StatusOK, updated)
}

func (h *CategoryHandler) ListCategories(w http.ResponseWriter, r *http.Request) {
	categories := h.service.GetCategories()
	helpers.RespondJSON(w, http.StatusOK, categories)
}

func (h *CategoryHandler) ListCategoriesByCategoryID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 32)
	if err != nil || id < 1 {
		helpers.RespondError(w, http.StatusBadRequest, "invalid category id")
		return
	}
	products, err := h.service.GetCategoriesByCategoryID(int(id))
	if err != nil {
		st := helpers.HTTPStatusForDB(err)
		if st == http.StatusNotFound {
			helpers.RespondError(w, http.StatusNotFound, "category not found")
			return
		}
	}
	helpers.RespondJSON(w, http.StatusOK, products)
}
