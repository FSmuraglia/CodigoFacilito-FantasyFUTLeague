package repositories

import (
	database "github.com/FSmuraglia/CodigoFacilito-FantasyFUTLeague/config"
	"github.com/FSmuraglia/CodigoFacilito-FantasyFUTLeague/internal/models"
	"gorm.io/gorm"
)

type TournamentRepository interface {
	GetAll(nameFilter string, sortParam string, status string) ([]models.Tournament, error)
	GetActiveTournamentsCount() (int64, error)
	GetTournamentWithTeamsAndMatches(tournamentID uint) (models.Tournament, []models.Match, error)
	GetTournamentsCountWonByTeamID(teamID uint) (int64, error)
}

type tournamentRepository struct {
	db *gorm.DB
}

func NewTournamentRepository() TournamentRepository {
	return &tournamentRepository{db: database.DB}
}

func (r *tournamentRepository) GetAll(nameFilter string, sortParam string, status string) ([]models.Tournament, error) {
	var tournaments []models.Tournament
	db := r.db

	if nameFilter != "" {
		db = db.Where("name LIKE ?", "%"+nameFilter+"%")
	}

	switch sortParam {
	case "prize_asc":
		db = db.Order("prize ASC")
	case "prize_desc":
		db = db.Order("prize DESC")
	case "date_asc":
		db = db.Order("start_date ASC")
	case "date_desc":
		db = db.Order("start_date DESC")
	}

	if status != "" {
		db = db.Where("status = ?", status)
	}

	err := db.Find(&tournaments).Error
	return tournaments, err
}

func (r *tournamentRepository) GetActiveTournamentsCount() (int64, error) {
	var count int64
	err := r.db.Model(&models.Tournament{}).Where("winner_id IS NULL").Count(&count).Error
	return count, err
}

func (r *tournamentRepository) GetTournamentWithTeamsAndMatches(tournamentID uint) (models.Tournament, []models.Match, error) {
	var tournament models.Tournament
	if err := r.db.Preload("Teams.Team").First(&tournament, tournamentID).Error; err != nil {
		return tournament, nil, err
	}

	var matches []models.Match
	if err := r.db.Where("tournament_id = ?", tournamentID).Find(&matches).Error; err != nil {
		return tournament, nil, err
	}

	return tournament, matches, nil
}

func (r *tournamentRepository) GetTournamentsCountWonByTeamID(teamID uint) (int64, error) {
	var count int64
	err := r.db.Model(&models.Tournament{}).Where("winner_id = ?", teamID).Count(&count).Error
	return count, err
}
