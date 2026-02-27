package handlers

import (
	"categories-api/models"
	"categories-api/services"
	"encoding/json"
	"net/http"
)

type TransactionHandler struct {
	service *services.TransactionService
}

func NewTransactionHandler(service *services.TransactionService) *TransactionHandler {
	return &TransactionHandler{service: service}
}

// multiple item apa saja , quantity
func (h *TransactionHandler) HandleCheckout(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		h.Checkout(w, r)
	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}

func (h *TransactionHandler) Checkout(w http.ResponseWriter, r *http.Request) {
	var req models.CheckoutRequest
	// 	type CheckoutItem struct {
	// 	ProductID int `json:"product_id"`
	// 	Quantity  int `json:"quantity"`
	// }
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	transaction, err := h.service.Checkout(req.Items)
// 	func (s *TransactionService) Checkout(items []models.CheckoutItem) (*models.Transaction, error) {
// 	return s.repo.CreateTransaction(items)
// }

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(transaction)
}
