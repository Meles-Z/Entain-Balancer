package dto

type TransactionRequest struct {
	State         string `json:"state"`
	Amount        string `json:"amount"`
	TransactionID string `json:"transactionId"`
}

type BalanceResponse struct {
	UserID  uint64 `json:"userId"`
	Balance string `json:"balance"`
}
