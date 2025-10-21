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

	var formattedTotalMarketValue string
	if lastFullTeam != nil {
		formattedTotalMarketValue = lastFullTeam.GetFormattedTotalMarketValue()
	}

	mostWinningTeam, _ := teamService.GetMostWinningTeam()
	upcomingMatches, _ := matchService.GetUpcomingMatches(3)

	utils.RenderTemplate(c, http.StatusOK, "index.html", gin.H{
		"ActiveTournaments":         activeTournaments,
		"TotalTeams":                totalTeams,
		"LastFullTeam":              lastFullTeam,
		"MostWinningTeam":           mostWinningTeam,
		"UpcomingMatches":           upcomingMatches,
		"FormattedTotalMarketValue": formattedTotalMarketValue,
	})
}
