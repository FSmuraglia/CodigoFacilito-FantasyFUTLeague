package models

type TournamentTeam struct {
	TournamentID uint
	TeamID       uint
	Team         Team `gorm:"foreignKey:TeamID"`
	Points       int  `gorm:"default:0;not null"`
}
