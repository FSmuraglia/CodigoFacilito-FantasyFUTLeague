package controllers

import (
	"fmt"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"time"

	database "github.com/FSmuraglia/CodigoFacilito-FantasyFUTLeague/config"
	log "github.com/FSmuraglia/CodigoFacilito-FantasyFUTLeague/internal/logger"
	"github.com/FSmuraglia/CodigoFacilito-FantasyFUTLeague/internal/models"
	"github.com/FSmuraglia/CodigoFacilito-FantasyFUTLeague/internal/services"
	"github.com/FSmuraglia/CodigoFacilito-FantasyFUTLeague/pkg/utils"
	"github.com/gin-gonic/gin"
)

var tournamentService *services.TournamentService

type TournamentWithFormattedPrize struct {
	models.Tournament
	FormattedPrize string
}

func InitTournamentController(s *services.TournamentService) {
	tournamentService = s
}

func CreateTournamentForm(c *gin.Context) {
	log.LogInfo("📝 Acceso a formulario de registro de torneo", nil)
	utils.RenderTemplate(c, http.StatusOK, "create_tournament.html", nil)
}

func CreateTournament(c *gin.Context) {
	var input struct {
		Name       string  `form:"name" binding:"required"`
		Prize      float64 `form:"prize" binding:"required"`
		StartDate  string  `form:"start_date" binding:"required"`
		EndDate    string  `form:"end_date"`
		TeamAmount int     `form:"team_amount" binding:"required,oneof=2 4"`
	}

	if err := c.ShouldBind(&input); err != nil {
		log.LogWarn("⚠️ Datos inválidos en creación del torneo", map[string]interface{}{
			"error":  err.Error(),
			"status": http.StatusBadRequest,
		})
		c.HTML(http.StatusBadRequest, "create_tournament.html", gin.H{
			"error": "Datos inválidos. Completá los datos correctamente",
		})
		return
	}

	startDate, err := time.Parse("2006-01-02", input.StartDate)
	if err != nil {
		log.LogWarn("⚠️ Fecha de inicio inválida", map[string]interface{}{
			"error":  err.Error(),
			"status": http.StatusBadRequest,
		})
		c.HTML(http.StatusBadRequest, "create_tournament.html", gin.H{
			"error": "Formado de fecha inválido",
		})
		return
	}

	var endDate *time.Time
	if input.EndDate != "" {
		parsedEndDate, err := time.Parse("2006-01-02", input.EndDate)
		if err == nil {
			endDate = &parsedEndDate
		}
	}

	tournament := models.Tournament{
		Name:       input.Name,
		Prize:      input.Prize,
		StartDate:  startDate,
		TeamAmount: input.TeamAmount,
	}

	if endDate != nil {
		tournament.EndDate = *endDate
	}

	if err := database.DB.Create(&tournament).Error; err != nil {
		log.LogError("❌ Error al crear torneo en la base de datos", map[string]interface{}{
			"error":  err.Error(),
			"status": http.StatusInternalServerError,
			"name":   input.Name,
		})
		c.HTML(http.StatusInternalServerError, "create_tournament.html", gin.H{
			"error": "Error al crear el torneo",
		})
		return
	}

	log.LogInfo("✅ Torneo creado correctamente", map[string]interface{}{
		"tournament_id": tournament.ID,
		"name":          tournament.Name,
		"status":        http.StatusSeeOther,
	})

	c.Redirect(http.StatusSeeOther, "/tournaments")
}

func ListTournaments(c *gin.Context) {
	nameFilter := strings.TrimSpace(c.Query("name"))
	sortParam := c.Query("sort")

	tournaments, err := tournamentService.ListTournaments(nameFilter, sortParam)
	if err != nil {
		log.LogError("❌ Error al obtener los torneos de la DB", map[string]interface{}{
			"status": http.StatusInternalServerError,
			"error":  err.Error(),
		})
		c.HTML(http.StatusInternalServerError, "tournaments.html", gin.H{
			"error": "Error al obtener los torneos de la DB",
		})
		return
	}

	var tournamentsFormatted []TournamentWithFormattedPrize
	for _, t := range tournaments {
		formatted := utils.FormatNumber(int64(t.Prize))
		tournamentsFormatted = append(tournamentsFormatted, TournamentWithFormattedPrize{
			Tournament:     t,
			FormattedPrize: formatted,
		})
	}

	isAdmin := false
	role, _ := utils.GetUserRoleFromCookie(c)
	if role == "ADMIN" {
		isAdmin = true
	}

	log.LogInfo("✅ Torneos obtenidos correctamente de la DB", map[string]interface{}{
		"count":  len(tournaments),
		"status": http.StatusOK,
	})
	utils.RenderTemplate(c, http.StatusOK, "tournaments.html", gin.H{
		"tournaments": tournamentsFormatted,
		"NameFilter":  nameFilter,
		"SortParam":   sortParam,
		"isAdmin":     isAdmin,
	})
}

func GetTournamentDetail(c *gin.Context) {

	// Struct para armar la tabla de posiciones en torneos de 4 equipos
	type TeamStats struct {
		TeamName     string
		BadgeURL     string
		Wins         int
		Losses       int
		GoalsFor     int
		GoalsAgainst int
		GoalDiff     int
		Points       int
	}

	id := c.Param("id")
	var tournament models.Tournament

	// Obtener el torneo
	if err := database.DB.
		Preload("Teams.Team").
		Preload("Winner").
		First(&tournament, id).Error; err != nil {
		log.LogError("❌ Torneo no encontrado", map[string]interface{}{
			"error":  err.Error(),
			"id":     id,
			"status": http.StatusNotFound,
		})
		c.HTML(http.StatusNotFound, "tournament_detail.html", gin.H{
			"error": "Torneo no encontrado",
		})
		return
	}

	// Formatear el Prize
	formattedPrize := utils.FormatNumber(int64(tournament.Prize))
	tournamentFormatted := TournamentWithFormattedPrize{
		Tournament:     tournament,
		FormattedPrize: formattedPrize,
	}

	// Obtener el usuario y verificar si está registrado en el torneo
	userID, exists := utils.GetUserIDFromCookie(c)

	isRegistered := false
	if exists {
		for _, t := range tournament.Teams {
			if t.Team.UserID == userID {
				isRegistered = true
				break
			}
		}
	}

	// Obtener la cantidad de equipos inscriptos para verificar si está lleno o no el torneo
	var teamCount int64
	database.DB.Model(&models.TournamentTeam{}).Where("tournament_id = ?", tournament.ID).Count(&teamCount)

	isFull := int(teamCount) >= tournament.TeamAmount

	if tournament.TeamAmount == 4 {
		var matches []models.Match
		if err := database.DB.Where("tournament_id = ? AND status = ?", tournament.ID, "FINISHED").Find(&matches).Error; err != nil {
			log.LogError("❌ Error al obtener los partidos del torneo", map[string]interface{}{
				"status": http.StatusInternalServerError,
				"error":  err.Error(),
			})
			c.HTML(http.StatusInternalServerError, "tournament_detail.html", gin.H{
				"error": "Error al obtener los partidos del torneo",
			})
			return
		}

		statsMap := make(map[uint]*TeamStats)

		for _, tt := range tournament.Teams {
			statsMap[tt.Team.ID] = &TeamStats{
				TeamName: tt.Team.Name,
				BadgeURL: tt.Team.BadgeUrl,
			}
		}

		for _, m := range matches {
			teamA := statsMap[m.TeamAID]
			teamB := statsMap[m.TeamBID]

			teamA.GoalsFor += m.TeamAGoals
			teamA.GoalsAgainst += m.TeamBGoals
			teamB.GoalsFor += m.TeamBGoals
			teamB.GoalsAgainst += m.TeamAGoals

			if m.TeamAGoals > m.TeamBGoals {
				teamA.Wins++
				teamB.Losses++
				teamA.Points += 3
			} else {
				teamB.Wins++
				teamA.Losses++
				teamB.Points += 3
			}
		}

		// Calcular diferencia de gol
		var standings []TeamStats
		for _, s := range statsMap {
			s.GoalDiff = s.GoalsFor - s.GoalsAgainst
			standings = append(standings, *s)
		}

		// Orden por puntos
		sort.Slice(standings, func(i, j int) bool {
			if standings[i].Points == standings[j].Points {
				return standings[i].GoalDiff > standings[j].GoalDiff
			}
			return standings[i].Points > standings[j].Points
		})
		utils.RenderTemplate(c, http.StatusOK, "tournament_detail.html", gin.H{
			"tournament":   tournamentFormatted,
			"isRegistered": isRegistered,
			"isFull":       isFull,
			"standings":    standings,
		})
	} else {
		utils.RenderTemplate(c, http.StatusOK, "tournament_detail.html", gin.H{
			"tournament":   tournamentFormatted,
			"isRegistered": isRegistered,
			"isFull":       isFull,
		})
	}
}

func JoinTournament(c *gin.Context) {

	userID, _ := utils.GetUserIDFromCookie(c)

	tournamentIDParam := c.Param("id")
	tournamentID, err := strconv.ParseUint(tournamentIDParam, 10, 64)
	if err != nil {
		log.LogError("❌ ID de torneo inválido", map[string]interface{}{
			"error":  err.Error(),
			"status": http.StatusBadRequest,
		})
		c.HTML(http.StatusBadRequest, "error.html", gin.H{
			"error": "ID de torneo inválido",
		})
		return
	}

	var team models.Team
	if err := database.DB.Where("user_id = ?", userID).First(&team).Error; err != nil {
		log.LogWarn("⚠️ El usuario no tiene un equipo creado", map[string]interface{}{
			"error":   err.Error(),
			"user_id": userID,
			"status":  http.StatusBadRequest,
		})
		c.HTML(http.StatusBadRequest, "error.html", gin.H{
			"error": "Debes crear un equipo antes de inscribirte a un torneo",
		})
		return
	}

	newRelation := models.TournamentTeam{
		TournamentID: uint(tournamentID),
		TeamID:       team.ID,
	}

	if err := database.DB.Create(&newRelation).Error; err != nil {
		log.LogError("❌ Error al inscribir equipo en torneo", map[string]interface{}{
			"error":  err.Error(),
			"status": http.StatusInternalServerError,
		})
		c.HTML(http.StatusInternalServerError, "error.html", gin.H{
			"error": "Error interno al inscribir al equipo",
		})
		return
	}

	log.LogInfo("✅ Equipo inscripto correctamente en el torneo", map[string]interface{}{
		"tournament_id": tournamentID,
		"team_id":       team.ID,
		"user_id":       userID,
	})

	c.Redirect(http.StatusSeeOther, fmt.Sprintf("/tournaments/%d", tournamentID))

}

func GetTeamsByTournament(c *gin.Context) {
	id := c.Param("id")
	var tournament models.Tournament

	if err := database.DB.Preload("Teams.Team").First(&tournament, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Torneo no encontrado",
		})
		return
	}

	var teams []gin.H
	for _, tt := range tournament.Teams {
		teams = append(teams, gin.H{
			"ID":   tt.Team.ID,
			"Name": tt.Team.Name,
		})
	}

	if len(teams) == 0 {
		c.JSON(http.StatusOK, []models.Team{})
		return
	}

	log.LogInfo("✅ Equipos obtenidos correctamente del torneo", map[string]interface{}{
		"tournament_id": id,
	})

	c.JSON(http.StatusOK, teams)

}
