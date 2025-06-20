package entities

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
	Model
	UserID     uint64           `gorm:"not null;index" json:"userId"`
	State      TransactionState `gorm:"type:varchar(10);not null" json:"state"`
	Amount     string           `gorm:"type:numeric(20,2);not null" json:"amount"`
	SourceType SourceType       `gorm:"type:varchar(20);not null" json:"sourceType"`
}
