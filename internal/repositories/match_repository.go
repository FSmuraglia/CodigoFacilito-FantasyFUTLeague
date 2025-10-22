package repositories

import (
	database "github.com/FSmuraglia/CodigoFacilito-FantasyFUTLeague/config"
	"github.com/FSmuraglia/CodigoFacilito-FantasyFUTLeague/internal/models"
	"gorm.io/gorm"
)

type TeamStatsProfile struct {
	MatchesPlayed int
	MatchesWon    int
	GoalsFor      int
	GoalsAgainst  int
	WinRate       float64
}

type MatchRepository interface {
	GetAll(sort string, status string) ([]models.Match, error)
	FindUpcomingMatches(limit int) ([]models.Match, error)
	CalculateTeamStats(teamID uint) (TeamStatsProfile, error)
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

func (r *matchRepository) CalculateTeamStats(teamID uint) (TeamStatsProfile, error) {
	var matches []models.Match
	err := r.db.Where("team_a_id = ? OR team_b_id = ?", teamID, teamID).
		Where("status = ?", "FINISHED").
		Find(&matches).Error

	var stats TeamStatsProfile

	if err != nil {
		return stats, err
	}

	for _, m := range matches {
		if m.TeamAID == teamID {
			stats.GoalsFor += m.TeamAGoals
			stats.GoalsAgainst += m.TeamBGoals
			if m.WinnerID != nil && *m.WinnerID == teamID {
				stats.MatchesWon++
			}
		} else {
			stats.GoalsFor += m.TeamBGoals
			stats.GoalsAgainst += m.TeamAGoals
			if m.WinnerID != nil && *m.WinnerID == teamID {
				stats.MatchesWon++
			}
		}
	}

	stats.MatchesPlayed = len(matches)
	if stats.MatchesPlayed > 0 {
		stats.WinRate = float64(stats.MatchesWon) / float64(stats.MatchesPlayed) * 100
	}

	return stats, nil
}
