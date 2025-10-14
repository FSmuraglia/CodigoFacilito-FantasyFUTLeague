package repositories

import (
	database "github.com/FSmuraglia/CodigoFacilito-FantasyFUTLeague/config"
	"github.com/FSmuraglia/CodigoFacilito-FantasyFUTLeague/internal/models"
	"gorm.io/gorm"
)

type TournamentRepository interface {
	GetAll(nameFilter string, sortParam string) ([]models.Tournament, error)
}

type tournamentRepository struct {
	db *gorm.DB
}

func NewTournamentRepository() TournamentRepository {
	return &tournamentRepository{db: database.DB}
}

func (r *tournamentRepository) GetAll(nameFilter string, sortParam string) ([]models.Tournament, error) {
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

	err := db.Find(&tournaments).Error
	return tournaments, err
}
