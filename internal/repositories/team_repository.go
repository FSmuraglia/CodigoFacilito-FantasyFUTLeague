package repositories

import (
	database "github.com/FSmuraglia/CodigoFacilito-FantasyFUTLeague/config"
	"github.com/FSmuraglia/CodigoFacilito-FantasyFUTLeague/internal/models"
	"gorm.io/gorm"
)

type TeamRepository interface {
	GetAll(nameFilter string, formationFilter string) ([]models.Team, error)
	GetTotalTeamsCount() (int64, error)
	FindLastCompleteTeam() (*models.Team, error)
	FindMostWinningTeam() (*models.Team, int64, error)
	GetByUserID(userID uint) (*models.Team, error)
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

func (r *teamRepository) GetTotalTeamsCount() (int64, error) {
	var count int64
	err := r.db.Model(&models.Team{}).Count(&count).Error
	return count, err
}

func (r *teamRepository) FindLastCompleteTeam() (*models.Team, error) {
	var team models.Team
	err := r.db.Preload("Players").Preload("User").
		Joins("JOIN players ON players.team_id = teams.id").
		Group("teams.id").
		Having("COUNT(players.id) = 11").
		Order("teams.id DESC").
		First(&team).Error
	return &team, err
}

func (r *teamRepository) FindMostWinningTeam() (*models.Team, int64, error) {
	// Struct para incluir los torneos ganados
	type TeamWinCount struct {
		ID   uint
		Wins int64
	}

	// Conseguir la cantidad de torneos ganados del equipo en la query
	var result TeamWinCount
	err := r.db.
		Table("teams").
		Select("teams.id, COUNT(tournaments.id) AS wins").
		Joins("JOIN tournaments ON tournaments.winner_id = teams.id").
		Group("teams.id").
		Order("wins DESC").
		Limit(1).
		Scan(&result).Error
	if err != nil {
		return nil, 0, err
	}

	// Conseguir los datos del equipo más ganador, para luego incluir en el index
	var team models.Team
	err = r.db.
		Preload("Players").
		Preload("User").
		First(&team, result.ID).Error
	if err != nil {
		return nil, 0, err
	}

	return &team, result.Wins, nil
}

func (r *teamRepository) GetByUserID(userID uint) (*models.Team, error) {
	var team models.Team
	err := r.db.Preload("Players").Where("user_id = ?", userID).First(&team).Error
	return &team, err
}
