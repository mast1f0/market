package handlers

import (
	"encoding/json"
	"fmt"
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

func (h *CategoryHandler) GetAllCategories(w http.ResponseWriter, r *http.Request) {
	categories := h.service.GetAll()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	categoriesJson, _ := json.Marshal(categories)
	w.Write(categoriesJson)
}

func (h *CategoryHandler) GetCategoryById(w http.ResponseWriter, r *http.Request) {
	categoryId, err := strconv.Atoi(chi.URLParam(r, "categoryId"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	category, err := h.service.Get(categoryId)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	categoryJson, _ := json.Marshal(category)
	w.Write(categoryJson)
}

func (h *CategoryHandler) DeleteCategoryById(w http.ResponseWriter, r *http.Request) {
	categoryId, err := strconv.Atoi(chi.URLParam(r, "categoryId"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = h.service.Delete(categoryId)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *CategoryHandler) AddCategory(w http.ResponseWriter, r *http.Request) {
	var Category *domain.Category
	err := json.NewDecoder(r.Body).Decode(&Category)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	CategoryJson, err := json.Marshal(Category)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(201)
	w.Write(CategoryJson)
}

func (h *CategoryHandler) Update(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("invalid id"))
	}
	var category *domain.Category
	category, err = h.service.Get(id)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("invalid id"))
		return
	}
	err = json.NewDecoder(r.Body).Decode(&category)
	if err != nil {
		http.Error(w, "Incorrect body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()
	categoryJson, err := json.Marshal(category)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("unable to get product"))
		return
	}
	w.Write(categoryJson)
}
