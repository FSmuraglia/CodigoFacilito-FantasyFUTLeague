package models

type TournamentTeam struct {
	TournamentID uint
	TeamID       uint
	Points       int `gorm:"default:0;not null"`
}
