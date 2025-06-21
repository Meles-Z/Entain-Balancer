package dto

type TransactionRequest struct {
	State         string `json:"state"`
	Amount        string `json:"amount"` // Keep as string for parsing later
	TransactionID string `json:"transactionId"`
}

type BalanceResponse struct {
	UserID  uint64 `json:"userId"`
	Balance string `json:"balance"` // Keep as string for consistent output
}
