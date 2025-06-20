package entities

type User struct {
	Model
	Name     string `gorm:"type:varchar(100);not null" json:"name"`
	Email    string `gorm:"type:varchar(100);uniqueIndex;not null" json:"email"`
	Password string `gorm:"type:varchar(255);not null" json:"password"`
	Balance  string `gorm:"type:numeric(20,2);not null" json:"balance"`
}
