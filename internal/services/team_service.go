package services

import (
	"github.com/FSmuraglia/CodigoFacilito-FantasyFUTLeague/internal/models"
	"github.com/FSmuraglia/CodigoFacilito-FantasyFUTLeague/internal/repositories"
)

type TeamService struct {
	repo repositories.TeamRepository
}

func NewTeamService(repo repositories.TeamRepository) *TeamService {
	return &TeamService{repo: repo}
}

func (s *TeamService) ListTeams(nameFilter, formationFilter string) ([]models.Team, error) {
	return s.repo.GetAll(nameFilter, formationFilter)
}
