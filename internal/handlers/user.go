package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/meles-z/entainbalancer/internal/dto"
)

func (h *Handler) GetUserBalance(w http.ResponseWriter, r *http.Request) {
	// Example path: /user/1/balance
	parts := strings.Split(strings.Trim(r.URL.Path, "/"), "/")
	if len(parts) != 3 || parts[0] != "user" || parts[2] != "balance" {
		http.Error(w, "Invalid path format", http.StatusBadRequest)
		return
	}

	userID, err := strconv.ParseUint(parts[1], 10, 64)
	if err != nil || userID == 0 {
		http.Error(w, "Invalid userId", http.StatusBadRequest)
		return
	}

	userBalance, err := h.UserService.GetUserByID(userID)
	if err != nil {
		http.Error(w, "Failed to get user balance", http.StatusInternalServerError)
		return
	}

	response := dto.BalanceResponse{
		UserID:  userID,
		Balance: userBalance.Balance.StringFixed(2),
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(response)
}
