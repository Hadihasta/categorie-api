package handlers

import (
	"categories-api/models"
	"categories-api/services"
	"encoding/json"
	"net/http"
)

type CategoryHandler struct {
	service *services.CategoryService
}


func NewCategoryHandler(service *services.CategoryService) *CategoryHandler{
	return &CategoryHandler{service : service}
}


func (h *CategoryHandler) HandleCategorys(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.GetAll(w, r)
	case http.MethodPost:
		h.Create(w, r)
	default: 
	http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}

func (h *CategoryHandler) GetAll(w http.ResponseWriter, r*http.Request){
	categories, err := h.service.GetAll()
	if err != nil{
		http.Error(w, err.Error(),http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(categories)
}


func (h *CategoryHandler) Create(w http.ResponseWriter, r*http.Request){
	var category models.Categories

	
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(&category); 
	err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	if category.Title == "" {
		http.Error(w, "title is required", http.StatusBadRequest)
		return
	}

	if err := h.service.Create(&category);
	 err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(category)
}

