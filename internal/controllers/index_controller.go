package controllers

import (
	"net/http"

	"github.com/FSmuraglia/CodigoFacilito-FantasyFUTLeague/pkg/utils"
	"github.com/gin-gonic/gin"
)

func ShowIndex(c *gin.Context) {
	activeTournaments, _ := tournamentService.GetActiveTournaments()
	totalTeams, _ := teamService.GetTotalTeams()

	lastFullTeam, _ := teamService.GetLastCompleteTeam()

	var lastFullTeamFormattedTotalMarketValue string
	if lastFullTeam != nil {
		lastFullTeamFormattedTotalMarketValue = lastFullTeam.GetFormattedTotalMarketValue()
	}

	mostWinningTeam, wins, _ := teamService.GetMostWinningTeam()

	var mostWinningTeamFormattedTotalMarketValue string
	if mostWinningTeam != nil {
		mostWinningTeamFormattedTotalMarketValue = mostWinningTeam.GetFormattedTotalMarketValue()
	}

	upcomingMatches, _ := matchService.GetUpcomingMatches(3)

	utils.RenderTemplate(c, http.StatusOK, "index.html", gin.H{
		"ActiveTournaments":                        activeTournaments,
		"TotalTeams":                               totalTeams,
		"LastFullTeam":                             lastFullTeam,
		"MostWinningTeam":                          mostWinningTeam,
		"MostWinningTeamWins":                      wins,
		"UpcomingMatches":                          upcomingMatches,
		"LastFullTeamFormattedTotalMarketValue":    lastFullTeamFormattedTotalMarketValue,
		"MostWinningTeamFormattedTotalMarketValue": mostWinningTeamFormattedTotalMarketValue,
	})
}
