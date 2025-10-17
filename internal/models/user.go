package models

type User struct {
	ID       uint    `gorm:"primarykey"`
	Username string  `gorm:"size:100;not null"`
	Password string  `gorm:"size:255;not null"`
	Email    string  `gorm:"size:255;uniqueIndex;not null"`
	Budget   float64 `gorm:"default:520000000"`
	Role     string  `gorm:"type:enum('USER','ADMIN');default:'USER'"`
}
