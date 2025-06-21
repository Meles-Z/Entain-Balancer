package entities

import (
	"time"

	"github.com/shopspring/decimal"
)

type User struct {
	ID        uint64          `gorm:"primaryKey;autoIncrement" json:"id"`
	Balance   decimal.Decimal `gorm:"type:numeric(20,2);not null" json:"balance"`
	CreatedAt time.Time       `json:"createdAt"`
	UpdatedAt time.Time       `json:"updatedAt"`
}
