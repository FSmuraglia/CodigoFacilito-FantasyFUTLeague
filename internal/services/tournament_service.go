package services

import (
	"github.com/FSmuraglia/CodigoFacilito-FantasyFUTLeague/internal/models"
	"github.com/FSmuraglia/CodigoFacilito-FantasyFUTLeague/internal/repositories"
)

type TournamentService struct {
	repo repositories.TournamentRepository
}

func NewTournamentService(repo repositories.TournamentRepository) *TournamentService {
	return &TournamentService{repo: repo}
}

func (s *TournamentService) ListTournaments(nameFilter, sortParam string) ([]models.Tournament, error) {
	return s.repo.GetAll(nameFilter, sortParam)
}
