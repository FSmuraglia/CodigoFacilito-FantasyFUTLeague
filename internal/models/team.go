package models

import "github.com/FSmuraglia/CodigoFacilito-FantasyFUTLeague/pkg/utils"

type Formation string

const (
	Formation433  Formation = "433"
	Formation4231 Formation = "4231"
	Formation442  Formation = "442"
)

type Team struct {
	ID             uint   `gorm:"primarykey"`
	Name           string `gorm:"size:100;not null"`
	UserID         uint
	User           User     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	BadgeUrl       string   `gorm:"size:255"`
	Players        []Player `gorm:"foreignKey:TeamID"`
	Tournaments    []TournamentTeam
	WonTournaments []Tournament `gorm:"foreignKey:WinnerID"`
	Formation      Formation    `gorm:"type:varchar(10);not null"`
}

func GetAvailableFormations() []Formation {
	return []Formation{Formation433, Formation4231, Formation442}
}

func (t *Team) CalculateRating() float64 {
	var total float64
	for _, player := range t.Players {
		total += player.Rating
	}
	return total / 11
}

func (t *Team) GetFormattedTotalMarketValue() string {
	var total float64
	for _, player := range t.Players {
		total += player.MarketValue
	}
	return utils.FormatNumber(int64(total))
}
