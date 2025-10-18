package controllers

import (
	"net/http"
	"strconv"
	"time"

	database "github.com/FSmuraglia/CodigoFacilito-FantasyFUTLeague/config"
	log "github.com/FSmuraglia/CodigoFacilito-FantasyFUTLeague/internal/logger"
	"github.com/FSmuraglia/CodigoFacilito-FantasyFUTLeague/internal/models"
	"github.com/FSmuraglia/CodigoFacilito-FantasyFUTLeague/internal/services"
	"github.com/FSmuraglia/CodigoFacilito-FantasyFUTLeague/pkg/utils"
	"github.com/gin-gonic/gin"
)

var matchService *services.MatchService

func InitMatchController(s *services.MatchService) {
	matchService = s
}

func ListMatches(c *gin.Context) {
	sort := c.Query("sort")
	status := c.Query("status")

	matches, err := matchService.ListMatches(sort, status)
	if err != nil {
		log.LogError("❌ Error al obtener los partidos de la DB", map[string]interface{}{
			"error":  err.Error(),
			"status": http.StatusInternalServerError,
		})
		c.HTML(http.StatusInternalServerError, "matches.html", gin.H{
			"error": "Error al obtener los partidos de la DB",
		})
		return
	}

	log.LogInfo("✅ Partidos obtenidos correctamente de la DB", map[string]interface{}{
		"count":  len(matches),
		"status": http.StatusOK,
	})

	isAdmin := false
	role, _ := utils.GetUserRoleFromCookie(c)
	if role == "ADMIN" {
		isAdmin = true
	}

	utils.RenderTemplate(c, http.StatusOK, "matches.html", gin.H{
		"matches": matches,
		"Sort":    sort,
		"Status":  status,
		"isAdmin": isAdmin,
	})
}

func CreateMatchForm(c *gin.Context) {
	log.LogInfo("📝 Acceso a formulario de registro de partido", nil)
	var tournaments []models.Tournament

	// Obtener torneos
	if err := database.DB.
		Preload("Teams").
		Find(&tournaments).Error; err != nil {
		log.LogError("❌ Error al obtener los torneos", map[string]interface{}{
			"error":  err.Error(),
			"status": http.StatusInternalServerError,
		})
		c.HTML(http.StatusInternalServerError, "create_match.html", gin.H{
			"error": "Error al cargar los torneos",
		})
		return
	}

	utils.RenderTemplate(c, http.StatusOK, "create_match.html", gin.H{
		"Tournaments": tournaments,
	})

}

func CreateMatch(c *gin.Context) {
	tournamentIDStr := c.PostForm("tournament_id")
	teamAIDStr := c.PostForm("team_a_id")
	teamBIDStr := c.PostForm("team_b_id")
	date := c.PostForm("date")

	tournamentID, _ := strconv.ParseUint(tournamentIDStr, 10, 64)
	teamAID, _ := strconv.ParseUint(teamAIDStr, 10, 64)
	teamBID, _ := strconv.ParseUint(teamBIDStr, 10, 64)

	if teamAID == teamBID {
		log.LogError("❌ El Admin intentó crear un partido entre un equipo y él mismo", nil)
		c.HTML(http.StatusBadRequest, "create_match.html", gin.H{
			"error": "Los equipos no pueden ser iguales",
		})
		return
	}

	parsedDate, _ := time.Parse("2006-01-02", date)

	match := models.Match{
		TournamentID: uint(tournamentID),
		TeamAID:      uint(teamAID),
		TeamBID:      uint(teamBID),
		Date:         parsedDate,
	}

	if err := database.DB.Create(&match).Error; err != nil {
		log.LogError("❌ Error al crear el partido", map[string]interface{}{
			"error":  err.Error(),
			"status": http.StatusInternalServerError,
		})
		c.HTML(http.StatusInternalServerError, "create_match.html", gin.H{
			"error": "Error al crear el partido",
		})
		return
	}

	log.LogInfo("✅ Partido creado correctamente", map[string]interface{}{
		"match_id": match.ID,
		"status":   http.StatusOK,
	})

	c.Redirect(http.StatusSeeOther, "/matches")

}

func GetMatchDetail(c *gin.Context) {
	id := c.Param("id")
	var match models.Match

	if err := database.DB.
		Preload("Tournament").
		Preload("TeamA.Players").
		Preload("TeamB.Players").
		First(&match, id).Error; err != nil {
		log.LogError("❌ Partido no encontrado", map[string]interface{}{
			"error":  err.Error(),
			"id":     id,
			"status": http.StatusNotFound,
		})
		c.HTML(http.StatusNotFound, "match_detail.html", gin.H{
			"error": "Partido no encontrado",
		})
		return
	}

	isAdmin := false
	role, _ := utils.GetUserRoleFromCookie(c)
	if role == "ADMIN" {
		isAdmin = true
	}

	utils.RenderTemplate(c, http.StatusOK, "match_detail.html", gin.H{
		"match":   match,
		"isAdmin": isAdmin,
	})
}
