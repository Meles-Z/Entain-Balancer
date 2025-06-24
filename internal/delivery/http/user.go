package http

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/meles-z/entainbalancer/internal/domain/user"
	"github.com/meles-z/entainbalancer/internal/infrastucture/logger"
)

func (h *Handler) GetUserBalance(w http.ResponseWriter, r *http.Request) {
	// Extract userId from URL path
	vars := mux.Vars(r)
	userIDStr := vars["userId"]

	userID, err := strconv.ParseUint(userIDStr, 10, 64)
	if err != nil || userID == 0 {
		logger.Warn("Invalid user ID in request", "userId", userIDStr)
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	// Fetch user balance from service
	userBalance, err := h.UserService.GetUserByID(userID)
	if err != nil {
		logger.Error("Failed to get user balance",
			"userId", userID,
			"error", err,
		)
		http.Error(w, "Failed to get user balance", http.StatusInternalServerError)
		return
	}

	// Prepare response
	response := user.BalanceResponse{
		UserID:    userID,
		Balance:   userBalance.Balance.StringFixed(2),
		CreatedAt: userBalance.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: userBalance.UpdatedAt.Format("2006-01-02 15:04:05"),
	}

	logger.Info("User balance fetched successfully",
		"userId", userID,
		"balance", response.Balance,
	)

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(response)
}
