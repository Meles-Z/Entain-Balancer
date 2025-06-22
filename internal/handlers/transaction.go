package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/meles-z/entainbalancer/internal/dto"
	"github.com/meles-z/entainbalancer/internal/entities"
	"github.com/meles-z/entainbalancer/internal/service"
)

type transactionHandler struct {
	service service.TransactionService
}

func NewTransactionHandler(service service.TransactionService) transactionHandler {
	return transactionHandler{
		service: service,
	}
}

func (h transactionHandler) UpdateTransaction(w http.ResponseWriter, r *http.Request) {
	// 1. Parse userId from the URL path: /user/{userId}/transaction
	// Example: "/user/123/transaction" → extract "123"
	parts := strings.Split(strings.Trim(r.URL.Path, "/"), "/")
	if len(parts) != 3 || parts[0] != "user" || parts[2] != "transaction" {
		http.Error(w, "Invalid URL format", http.StatusBadRequest)
		return
	}

	userID, err := strconv.ParseUint(parts[1], 10, 64)
	if err != nil || userID == 0 {
		http.Error(w, "Invalid userId", http.StatusBadRequest)
		return
	}

	// 2. Parse Source-Type from header
	sourceType := r.Header.Get("Source-Type")
	if sourceType == "" {
		http.Error(w, "Missing Source-Type header", http.StatusBadRequest)
		return
	}
	if sourceType != "game" && sourceType != "server" && sourceType != "payment" {
		http.Error(w, "Invalid Source-Type header", http.StatusBadRequest)
		return
	}

	// 3. Decode body
	var req dto.TransactionRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// 4. Validate state
	if req.State != string(entities.TransactionStateWin) && req.State != string(entities.TransactionStateLose) {
		http.Error(w, "Invalid transaction state", http.StatusBadRequest)
		return
	}

	// 5. Call service
	transaction := &entities.Transaction{
		TransactionID: req.TransactionID,
		UserID:        userID,
		State:         entities.TransactionState(req.State),
		Amount:        req.Amount,
		SourceType:    entities.SourceType(sourceType),
	}

	err = h.service.UpdateTransaction(transaction)
	if err != nil {
		switch {
		case errors.Is(err, service.ErrTransactionAlreadyProcessed):
			http.Error(w, "Transaction already processed", http.StatusConflict)
		case errors.Is(err, service.ErrInsufficientBalance):
			http.Error(w, "Insufficient balance", http.StatusBadRequest)
		default:
			http.Error(w, "Internal server error: "+err.Error(), http.StatusInternalServerError)
		}
		return
	}

	// 6. Response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Transaction processed successfully",
	})
}
