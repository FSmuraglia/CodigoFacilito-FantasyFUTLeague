package models

import "time"

type Status string

const (
	StatusNotStarted Status = "NOT STARTED"
	StatusInProgress Status = "IN PROGRESS"
	StatusFinished   Status = "FINISHED"
)

type Match struct {
	ID           uint       `gorm:"primarykey"`
	TournamentID uint       `gorm:"not null"`
	Tournament   Tournament `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	TeamAID      uint
	TeamA        Team `gorm:"foreignKey:TeamAID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	TeamBID      uint
	TeamB        Team      `gorm:"foreignKey:TeamBID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Date         time.Time `gorm:"type:date;not null"`
	Status       Status    `gorm:"type:varchar(20);default:'NOT STARTED'"`
	WinnerID     *uint
	Winner       *Team `gorm:"foreignKey:WinnerID"`
	TeamAGoals   int   `gorm:"default:0;not null"`
	TeamBGoals   int   `gorm:"default:0;not null"`
}
