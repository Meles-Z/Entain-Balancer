package user

type BalanceResponse struct {
	UserID    uint64 `json:"userId"`
	Balance   string `json:"balance"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}