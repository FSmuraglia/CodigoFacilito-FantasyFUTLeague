package services

import (
	"sort"

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

func (s *TournamentService) GetActiveTournaments() (int64, error) {
	return s.repo.GetActiveTournamentsCount()
}

func (s *TournamentService) GetTournamentsWonByTeamID(teamID uint) (int64, error) {
	return s.repo.GetTournamentsCountWonByTeamID(teamID)
}

func (s *TournamentService) CalculateTournamentTable(tournamentID uint) ([]models.TeamStats, error) {
	tournament, matches, err := s.repo.GetTournamentWithTeamsAndMatches(tournamentID)
	if err != nil {
		return nil, err
	}

	stats := make(map[uint]*models.TeamStats)

	// Inicializar equipos
	for _, tt := range tournament.Teams {
		stats[tt.Team.ID] = &models.TeamStats{
			TeamID:   tt.Team.ID,
			TeamName: tt.Team.Name,
			BadgeURL: tt.Team.BadgeUrl,
		}
	}

	// Calcular estadísticas desde los partidos finalizados
	for _, match := range matches {
		if match.Status != "FINISHED" {
			continue
		}

		teamA := stats[match.TeamAID]
		teamB := stats[match.TeamBID]
		if teamA == nil || teamB == nil {
			continue
		}

		// Goles a favor y en contra
		teamA.GoalsFor += match.TeamAGoals
		teamA.GoalsAgainst += match.TeamBGoals
		teamB.GoalsFor += match.TeamBGoals
		teamB.GoalsAgainst += match.TeamAGoals

		// Diferencia de gol
		teamA.GoalDifference = teamA.GoalsFor - teamA.GoalsAgainst
		teamB.GoalDifference = teamB.GoalsFor - teamB.GoalsAgainst

		// Resultado
		if match.TeamAGoals > match.TeamBGoals {
			teamA.Wins++
			teamA.Points += 3
			teamB.Losses++
		} else if match.TeamBGoals > match.TeamAGoals {
			teamB.Wins++
			teamB.Points += 3
			teamA.Losses++
		}
	}

	// Convertir a slice
	var table []models.TeamStats
	for _, t := range stats {
		table = append(table, *t)
	}

	// Ordenar tabla por puntos y diferencia de gol
	sort.Slice(table, func(i, j int) bool {
		if table[i].Points == table[j].Points {
			return table[i].GoalDifference > table[j].GoalDifference
		}
		return table[i].Points > table[j].Points
	})

	return table, nil
}
