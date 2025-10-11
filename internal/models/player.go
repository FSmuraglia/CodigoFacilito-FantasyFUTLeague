package models

type Player struct {
	ID          uint `gorm:"primarykey"`
	TeamID      uint
	Team        Team    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
	Name        string  `gorm:"size:255;not null"`
	Nationality string  `gorm:"size:100;not null"`
	MarketValue float64 `gorm:"not null"`
	PhotoUrl    string  `gorm:"size:255"`
}
