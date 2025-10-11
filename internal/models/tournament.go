package models

import "time"

type Tournament struct {
	ID        uint `gorm:"primarykey"`
	Teams     []TournamentTeam
	Name      string `gorm:"size:255;not null"`
	Prize     float64
	StartDate time.Time `gorm:"type:date;not null"`
	EndDate   time.Time `gorm:"type:date"`
	WinnerID  uint
	Winner    Team    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
	Matches   []Match `gorm:"foreignKey:TournamentID"`
}
