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

func (h *Handler) UpdateTransaction(w http.ResponseWriter, r *http.Request) {
	// 1. Parse and validate URL path
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

	// 2. Validate Source-Type header
	sourceType := r.Header.Get("Source-Type")
	if sourceType != "game" && sourceType != "server" && sourceType != "payment" {
		http.Error(w, "Invalid or missing Source-Type header", http.StatusBadRequest)
		return
	}

	// 3. Decode and validate request body
	var req dto.TransactionRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if req.State != string(entities.TransactionStateWin) && req.State != string(entities.TransactionStateLose) {
		http.Error(w, "Invalid transaction state", http.StatusBadRequest)
		return
	}

	// 4. Build transaction entity
	transaction := &entities.Transaction{
		TransactionID: req.TransactionID,
		UserID:        userID,
		State:         entities.TransactionState(req.State),
		Amount:        req.Amount,
		SourceType:    entities.SourceType(sourceType),
	}

	// 5. Process transaction via service
	err = h.TransactionService.UpdateTransaction(transaction)
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

	// 6. Success response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(map[string]string{
		"message": "Transaction processed successfully",
	})
}
