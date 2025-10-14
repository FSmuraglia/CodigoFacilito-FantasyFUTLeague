package repositories

import (
	database "github.com/FSmuraglia/CodigoFacilito-FantasyFUTLeague/config"
	"github.com/FSmuraglia/CodigoFacilito-FantasyFUTLeague/internal/models"
	"gorm.io/gorm"
)

type TeamRepository interface {
	GetAll(nameFilter string, formationFilter string) ([]models.Team, error)
}

type teamRepository struct {
	db *gorm.DB
}

func NewTeamRepository() TeamRepository {
	return &teamRepository{db: database.DB}
}

func (r *teamRepository) GetAll(nameFilter string, formationFilter string) ([]models.Team, error) {
	var teams []models.Team
	db := r.db.Preload("User")

	if nameFilter != "" {
		db = db.Where("name LIKE ?", "%"+nameFilter+"%")
	}
	if formationFilter != "" {
		db = db.Where("formation = ?", formationFilter)
	}

	err := db.Find(&teams).Error
	return teams, err
}
