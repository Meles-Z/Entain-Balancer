package entities

import (
	"time"
)

type User struct {
	ID        uint64    `gorm:"primaryKey" json:"userId"`
	Balance   float64   `gorm:"type:numeric(20,2);not null" json:"balance"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
