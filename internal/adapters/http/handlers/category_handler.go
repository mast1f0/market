package handlers

import (
	"encoding/json"
	"market/internal/core/domain"
	"market/internal/core/service"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type CategoryHandler struct {
	service *service.CategoryService
}

func NewCategoryHandler(service *service.CategoryService) *CategoryHandler {
	return &CategoryHandler{service: service}
}

func (h *CategoryHandler) AddCategory(w http.ResponseWriter, r *http.Request) {
	var category domain.Category
	err := json.NewDecoder(r.Body).Decode(&category)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()
	_, err = h.service.CreateCategory(&category)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	jsonCategory, _ := json.Marshal(category)
	w.Write(jsonCategory)
}

func (h *CategoryHandler) GetCategory(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("invalid category id"))
		return
	}
	category, err := h.service.GetCategory(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	categoryJSON, _ := json.Marshal(category)
	w.WriteHeader(http.StatusOK)
	w.Write(categoryJSON)
}

func (h *CategoryHandler) DeleteCategory(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("invalid category id"))
	}
	err = h.service.DeleteCategory(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *CategoryHandler) UpdateCategory(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("invalid category id"))
		return
	}
	var category = domain.Category{
		ID: uint(id),
	}
	err = json.NewDecoder(r.Body).Decode(&category)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	defer r.Body.Close()
	newCategory, err := h.service.UpdateCategory(&category)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	jsonCategory, _ := json.Marshal(newCategory)
	w.Write(jsonCategory)
}
