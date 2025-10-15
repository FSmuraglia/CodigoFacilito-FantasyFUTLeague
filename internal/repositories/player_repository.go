package repositories

import (
	database "github.com/FSmuraglia/CodigoFacilito-FantasyFUTLeague/config"
	"github.com/FSmuraglia/CodigoFacilito-FantasyFUTLeague/internal/models"
	"gorm.io/gorm"
)

type PlayerRepository interface {
	GetAll(nameFilter string, positionFilter string, sortParam string) ([]models.Player, error)
}

type playerRepository struct {
	db *gorm.DB
}

func NewPlayerRepository() PlayerRepository {
	return &playerRepository{db: database.DB}
}

func (r *playerRepository) GetAll(nameFilter string, positionFilter string, sortParam string) ([]models.Player, error) {
	var players []models.Player
	db := r.db.Preload("Team")

	if nameFilter != "" {
		db = db.Where("name LIKE ?", "%"+nameFilter+"%")
	}

	if positionFilter != "" {
		db = db.Where("position = ?", positionFilter)
	}

	switch sortParam {
	case "value_desc":
		db = db.Order("market_value DESC")
	case "value_asc":
		db = db.Order("market_value ASC")
	}

	err := db.Find(&players).Error

	return players, err
}
