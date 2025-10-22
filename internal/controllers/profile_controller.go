package controllers

import (
	"net/http"

	log "github.com/FSmuraglia/CodigoFacilito-FantasyFUTLeague/internal/logger"
	"github.com/FSmuraglia/CodigoFacilito-FantasyFUTLeague/internal/repositories"
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

	isAdmin := false
	role, _ := utils.GetUserRoleFromCookie(c)
	if role == "ADMIN" {
		isAdmin = true
	}

	user := userService.GetUser(userID)
	formattedBudget := utils.FormatNumber(int64(user.Budget))

	team, err := teamService.GetTeamByUserID(userID)
	if err != nil {
		log.LogError("❌ Error al obtener el equipo del usuario", map[string]interface{}{
			"error":  err.Error(),
			"status": http.StatusInternalServerError,
		})
	}

	var stats repositories.TeamStatsProfile
	var tournamentsWon int64

	if team != nil {
		stats, _ = matchService.GetTeamStats(team.ID)
		tournamentsWon, _ = tournamentService.GetTournamentsWonByTeamID(team.ID)
	}

	utils.RenderTemplate(c, http.StatusOK, "profile.html", gin.H{
		"user":            user,
		"team":            team,
		"stats":           stats,
		"tournamentsWon":  tournamentsWon,
		"formattedBudget": formattedBudget,
		"isAdmin":         isAdmin,
	})

}
