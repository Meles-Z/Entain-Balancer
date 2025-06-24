package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/meles-z/entainbalancer/internal/dto"
)

func (h *Handler) GetUserBalance(w http.ResponseWriter, r *http.Request) {
	// Extract userId from URL path
	vars := mux.Vars(r)
	userIDStr := vars["userId"]

	userID, err := strconv.ParseUint(userIDStr, 10, 64)
	if err != nil || userID == 0 {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}
	// Fetch user balance from service
	userBalance, err := h.UserService.GetUserByID(userID)
	if err != nil {
		log.Printf("Failed to get user balance for ID %d: %v\n", userID, err)
		http.Error(w, "Failed to get user balance", http.StatusInternalServerError)
		return
	}

	// Prepare response
	response := dto.BalanceResponse{
		UserID:  userID,
		Balance: userBalance.Balance.StringFixed(2),
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(response)
}
