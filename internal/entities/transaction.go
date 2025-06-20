package entities

import (
	"time"
)

type TransactionState string

const (
	TransactionStateWin  TransactionState = "win"
	TransactionStateLose TransactionState = "lose"
)

type SourceType string

const (
	SourceTypeGame    SourceType = "game"
	SourceTypeServer  SourceType = "server"
	SourceTypePayment SourceType = "payment"
)

type Transaction struct {
	ID            uint             `gorm:"primaryKey"`
	TransactionID string           `gorm:"uniqueIndex;not null" json:"transactionId"`
	UserID        uint64           `gorm:"not null;index" json:"userId"`
	State         TransactionState `gorm:"type:varchar(10);not null" json:"state"`
	Amount        float64          `gorm:"type:numeric(20,2);not null" json:"amount"`
	SourceType    SourceType       `gorm:"type:varchar(20);not null" json:"sourceType"`
	CreatedAt     time.Time        `json:"created_at"`
}
