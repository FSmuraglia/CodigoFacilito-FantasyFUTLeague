package repositories

import (
	database "github.com/FSmuraglia/CodigoFacilito-FantasyFUTLeague/config"
	"github.com/FSmuraglia/CodigoFacilito-FantasyFUTLeague/internal/models"
	"gorm.io/gorm"
)

type MatchRepository interface {
	GetAll(sort string, status string) ([]models.Match, error)
	FindUpcomingMatches(limit int) ([]models.Match, error)
}

type matchRepository struct {
	db *gorm.DB
}

func NewMatchRepository() MatchRepository {
	return &matchRepository{db: database.DB}
}

func (r *matchRepository) GetAll(sort string, status string) ([]models.Match, error) {
	var matches []models.Match
	query := r.db.Preload("Tournament").Preload("TeamA").Preload("TeamB")

	if status != "" {
		query = query.Where("status = ?", status)
	}

	// Orden por fecha
	if sort == "asc" {
		query = query.Order("date ASC")
	} else if sort == "desc" {
		query = query.Order("date DESC")
	}

	err := query.Find(&matches).Error
	return matches, err

}

func (r *matchRepository) FindUpcomingMatches(limit int) ([]models.Match, error) {
	var matches []models.Match
	err := r.db.Preload("TeamA").Preload("TeamB").Preload("Tournament").
		Where("status = ?", "NOT STARTED").
		Order("date ASC").
		Limit(limit).
		Find(&matches).Error
	return matches, err
}
