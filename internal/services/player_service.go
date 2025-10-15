package services

import (
	"github.com/FSmuraglia/CodigoFacilito-FantasyFUTLeague/internal/models"
	"github.com/FSmuraglia/CodigoFacilito-FantasyFUTLeague/internal/repositories"
)

type PlayerService struct {
	repo repositories.PlayerRepository
}

func NewPlayerService(repo repositories.PlayerRepository) *PlayerService {
	return &PlayerService{repo: repo}
}

func (s *PlayerService) ListPlayers(nameFilter, positionFilter, sortParam string) ([]models.Player, error) {
	return s.repo.GetAll(nameFilter, positionFilter, sortParam)
}
