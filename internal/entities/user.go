package entities

import "github.com/shopspring/decimal"

type User struct {
	Model
	Balance decimal.Decimal `gorm:"type:numeric(20,2);not null" json:"balance"`
}
