package http

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/meles-z/entainbalancer/internal/domain/transaction"
	"github.com/meles-z/entainbalancer/internal/infrastructure/logger"
)

func (h *Handler) UpdateTransaction(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userIDStr := vars["userId"]

	userID, err := strconv.ParseUint(userIDStr, 10, 64)
	if err != nil || userID == 0 {
		logger.Warn("Invalid user ID in request", "userId", userIDStr)
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	sourceType := r.Header.Get("Source-Type")
	if sourceType != "game" && sourceType != "server" && sourceType != "payment" {
		logger.Warn("Invalid or missing Source-Type header", "sourceType", sourceType)
		http.Error(w, "Invalid or missing Source-Type header", http.StatusBadRequest)
		return
	}

	var req transaction.TransactionRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		logger.Warn("Failed to decode request body", "error", err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if req.State != string(transaction.TransactionStateWin) && req.State != string(transaction.TransactionStateLose) {
		logger.Warn("Invalid transaction state", "state", req.State)
		http.Error(w, "Invalid transaction state", http.StatusBadRequest)
		return
	}

	newTxn := &transaction.Transaction{
		TransactionID: req.TransactionID,
		UserID:        userID,
		State:         transaction.TransactionState(req.State),
		Amount:        req.Amount,
		SourceType:    transaction.SourceType(sourceType),
	}

	err = h.TransactionService.UpdateTransaction(newTxn)
	if err != nil {
		switch {
		case errors.Is(err, transaction.ErrTransactionAlreadyProcessed):
			logger.Warn("Transaction already processed", "transactionID", req.TransactionID)
			http.Error(w, "Transaction already processed", http.StatusConflict)
		case errors.Is(err, transaction.ErrInsufficientBalance):
			logger.Warn("Insufficient balance for transaction", "transactionID", req.TransactionID)
			http.Error(w, "Insufficient balance", http.StatusBadRequest)
		default:
			logger.Error("Unexpected error processing transaction",
				"transactionID", req.TransactionID,
				"error", err)
			http.Error(w, "Internal server error: "+err.Error(), http.StatusInternalServerError)
		}
		return
	}

	logger.Info("Transaction processed successfully",
		"userID", userID,
		"transactionID", req.TransactionID,
		"amount", req.Amount,
		"state", req.State,
		"sourceType", sourceType,
	)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(map[string]string{
		"message": "Transaction processed successfully",
	})
}
