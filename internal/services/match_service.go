package services

import (
	"github.com/FSmuraglia/CodigoFacilito-FantasyFUTLeague/internal/models"
	"github.com/FSmuraglia/CodigoFacilito-FantasyFUTLeague/internal/repositories"
)

type MatchService struct {
	repo repositories.MatchRepository
}

func NewMatchService(repo repositories.MatchRepository) *MatchService {
	return &MatchService{repo: repo}
}

func (s *MatchService) ListMatches(sort string, status string) ([]models.Match, error) {
	return s.repo.GetAll(sort, status)
}

func (s *MatchService) GetUpcomingMatches(limit int) ([]models.Match, error) {
	return s.repo.FindUpcomingMatches(limit)
}

func (s *MatchService) GetTeamStats(teamID uint) (repositories.TeamStatsProfile, error) {
	return s.repo.CalculateTeamStats(teamID)
}
