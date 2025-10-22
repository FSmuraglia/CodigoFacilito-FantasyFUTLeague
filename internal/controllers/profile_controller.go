package controllers

import (
	"net/http"

	"github.com/FSmuraglia/CodigoFacilito-FantasyFUTLeague/internal/services"
	"github.com/FSmuraglia/CodigoFacilito-FantasyFUTLeague/pkg/utils"
	"github.com/gin-gonic/gin"
)

var userService *services.UserService

func InitProfileController(s *services.UserService) {
	userService = s
}

func GetProfile(c *gin.Context) {
	userID, exists := utils.GetUserIDFromCookie(c)
	if !exists {
		c.Redirect(http.StatusSeeOther, "/login")
		return
	}

	user := userService.GetUser(userID)
	formattedBudget := utils.FormatNumber(int64(user.Budget))

	team, _ := teamService.GetTeamByUserID(userID)

	stats, _ := matchService.GetTeamStats(team.ID)
	tournamentsWon, _ := tournamentService.GetTournamentsWonByTeamID(team.ID)

	utils.RenderTemplate(c, http.StatusOK, "profile.html", gin.H{
		"user":            user,
		"team":            team,
		"stats":           stats,
		"tournamentsWon":  tournamentsWon,
		"formattedBudget": formattedBudget,
	})

}
