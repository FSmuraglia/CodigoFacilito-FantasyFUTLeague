package controllers

import (
	"fmt"
	"net/http"

	database "github.com/FSmuraglia/CodigoFacilito-FantasyFUTLeague/config"
	log "github.com/FSmuraglia/CodigoFacilito-FantasyFUTLeague/internal/logger"
	"github.com/FSmuraglia/CodigoFacilito-FantasyFUTLeague/internal/models"
	"github.com/FSmuraglia/CodigoFacilito-FantasyFUTLeague/internal/services"
	"github.com/FSmuraglia/CodigoFacilito-FantasyFUTLeague/pkg/utils"
	"github.com/gin-gonic/gin"
)

var playerService *services.PlayerService

func InitPlayerController(s *services.PlayerService) {
	playerService = s
}

func CreatePlayerForm(c *gin.Context) {
	log.LogInfo("📝 Acceso a formulario de registro de jugador", nil)
	utils.RenderTemplate(c, http.StatusOK, "create_player.html", gin.H{
		"Positions": models.GetAvailablePositions(),
	})
}

func CreatePlayer(c *gin.Context) {
	var input struct {
		Name        string  `form:"name" binding:"required"`
		Nationality string  `form:"nationality" binding:"required"`
		MarketValue float64 `form:"market_value" binding:"required"`
		Rating      float64 `form:"rating" binding:"required"`
		PhotoUrl    string  `form:"photo_url" binding:"required"`
		Position    string  `form:"position" binding:"required"`
	}

	if err := c.ShouldBind(&input); err != nil {
		log.LogWarn("⚠️ Datos inválidos al crear el jugador", map[string]interface{}{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
		c.HTML(http.StatusBadRequest, "create_player.html", gin.H{
			"error":     "Datos inválidos al crear el jugador",
			"Positions": models.GetAvailablePositions(),
		})
		return
	}

	player := models.Player{
		Name:        input.Name,
		Nationality: input.Nationality,
		MarketValue: input.MarketValue,
		Rating:      input.Rating,
		PhotoUrl:    input.PhotoUrl,
		Position:    models.Position(input.Position),
	}

	if err := database.DB.Create(&player).Error; err != nil {
		log.LogError("❌ Error al crear el jugador en la DB", map[string]interface{}{
			"error":  err.Error(),
			"status": http.StatusInternalServerError,
		})
		c.HTML(http.StatusInternalServerError, "create_player.html", gin.H{
			"error":     "Error al crear el jugador",
			"Positions": models.GetAvailablePositions(),
		})
		return
	}

	log.LogInfo("✅ Jugador creado correctamente", map[string]interface{}{
		"status": http.StatusSeeOther,
		"player": player.Name,
	})

	c.Redirect(http.StatusSeeOther, "/players")

}

func ListPlayers(c *gin.Context) {
	nameFilter := c.Query("name")
	positionFilter := c.Query("position")
	sortParam := c.Query("sort")

	players, err := playerService.ListPlayers(nameFilter, positionFilter, sortParam)
	if err != nil {
		log.LogError("❌ Error al obtener los jugadores de la DB", map[string]interface{}{
			"error":  err.Error(),
			"status": http.StatusInternalServerError,
		})
		c.HTML(http.StatusInternalServerError, "players.html", gin.H{
			"error": "Error al obtener los jugadores de la DB",
		})
		return
	}

	type PlayerWithFormattedValue struct {
		models.Player
		FormattedValue string
	}
	var playersFormatted []PlayerWithFormattedValue
	for _, p := range players {
		formatted := utils.FormatNumber(int64(p.MarketValue))
		playersFormatted = append(playersFormatted, PlayerWithFormattedValue{
			Player:         p,
			FormattedValue: formatted,
		})
	}

	isAdmin := false
	role, _ := utils.GetUserRoleFromCookie(c)
	if role == "ADMIN" {
		isAdmin = true
	}

	log.LogInfo("✅ Jugadores obtenidos correctamente de la DB", map[string]interface{}{
		"count":  len(players),
		"status": http.StatusOK,
	})

	utils.RenderTemplate(c, http.StatusOK, "players.html", gin.H{
		"Players":   playersFormatted,
		"Positions": models.GetAvailablePositions(),
		"isAdmin":   isAdmin,
	})
}

func BuyPlayer(c *gin.Context) {
	userID, _ := utils.GetUserIDFromCookie(c)
	playerID := c.Param("id")

	nameFilter := c.Query("name")
	positionFilter := c.Query("position")
	sortParam := c.Query("sort")

	players, _ := playerService.ListPlayers(nameFilter, positionFilter, sortParam)

	type PlayerWithFormattedValue struct {
		models.Player
		FormattedValue string
	}
	var playersFormatted []PlayerWithFormattedValue
	for _, p := range players {
		formatted := utils.FormatNumber(int64(p.MarketValue))
		playersFormatted = append(playersFormatted, PlayerWithFormattedValue{
			Player:         p,
			FormattedValue: formatted,
		})
	}

	// Buscar jugador
	var player models.Player
	if err := database.DB.First(&player, playerID).Error; err != nil {
		log.LogError("❌ Error al obtener el jugador de la DB", map[string]interface{}{
			"error":  err.Error(),
			"status": http.StatusNotFound,
		})
		utils.RenderTemplate(c, http.StatusNotFound, "players.html", gin.H{
			"error":     "Jugador no encontrado",
			"Players":   playersFormatted,
			"Positions": models.GetAvailablePositions(),
		})
		return
	}

	// Buscar equipo del usuario
	var team models.Team
	if err := database.DB.Preload("User").Where("user_id = ?", userID).First(&team).Error; err != nil {
		log.LogError("❌ El usuario intentó comprar un jugador sin tener un equipo creado", map[string]interface{}{
			"error":  err.Error(),
			"status": http.StatusBadRequest,
		})
		utils.RenderTemplate(c, http.StatusBadRequest, "players.html", gin.H{
			"error":     "Debes crear un equipo antes de poder comprar jugadores",
			"Players":   playersFormatted,
			"Positions": models.GetAvailablePositions(),
		})
		return
	}

	// Obtener distribución de la formación y verificar jugadores en esa posición
	requirements := utils.GetFormationRequirements(string(team.Formation))
	var count int64
	database.DB.Model(&models.Player{}).
		Where("team_id = ? AND position = ?", team.ID, player.Position).
		Count(&count)

	maxAllowed, ok := requirements[string(player.Position)]
	if !ok {
		log.LogError("❌ El usuario intentó comprar un jugador cuya posición no encaja en su formación de equipo", map[string]interface{}{
			"status": http.StatusBadRequest,
			"player": player.Name,
		})
		utils.RenderTemplate(c, http.StatusBadRequest, "players.html", gin.H{
			"error":     fmt.Sprintf("Tu formación no admite jugadores en la posición %s", player.Position),
			"Players":   playersFormatted,
			"Positions": models.GetAvailablePositions(),
		})
		return
	}

	if int(count) >= maxAllowed {
		log.LogError("❌ El usuario intentó comprar un jugador cuya posición ya tiene cubierta en su equipo", map[string]interface{}{
			"status": http.StatusBadRequest,
			"player": player.Name,
		})
		utils.RenderTemplate(c, http.StatusBadRequest, "players.html", gin.H{
			"error":     fmt.Sprintf("Ya tienes el máximo permitido de jugadores en la posición %s", player.Position),
			"Players":   playersFormatted,
			"Positions": models.GetAvailablePositions(),
		})
		return
	}

	// Verificar presupuesto
	if team.User.Budget < player.MarketValue {
		log.LogError("❌ El usuario no tiene el presupuesto suficiente para comprar al jugador", map[string]interface{}{
			"status": http.StatusBadRequest,
			"player": player.Name,
		})
		utils.RenderTemplate(c, http.StatusBadRequest, "players.html", gin.H{
			"error":     "No tienes suficiente presupuesto para comprar a este jugador",
			"Players":   playersFormatted,
			"Positions": models.GetAvailablePositions(),
		})
		return
	}

	// Asignar jugador al equipo y descontar presupuesto
	player.TeamID = &team.ID
	if err := database.DB.Save(&player).Error; err != nil {
		log.LogError("❌ Error al asignar el jugador al equipo del usuario", map[string]interface{}{
			"status":  http.StatusInternalServerError,
			"player":  player.Name,
			"user_id": userID,
		})
		utils.RenderTemplate(c, http.StatusInternalServerError, "players.html", gin.H{
			"error":     "Error al asignar jugador al equipo",
			"Players":   playersFormatted,
			"Positions": models.GetAvailablePositions(),
		})
		return
	}

	team.User.Budget -= player.MarketValue
	if err := database.DB.Save(&team.User).Error; err != nil {
		log.LogError("❌ Error al actualizar el presupuesto del usuario", map[string]interface{}{
			"status":  http.StatusInternalServerError,
			"user_id": userID,
		})
		utils.RenderTemplate(c, http.StatusInternalServerError, "players.html", gin.H{
			"error":     "Error al actualizar el presupuesto del usuario",
			"Players":   playersFormatted,
			"Positions": models.GetAvailablePositions(),
		})
		return
	}

	log.LogInfo("✅ Jugador comprado correctamente", map[string]interface{}{
		"user_id": userID,
		"player":  player.Name,
		"status":  http.StatusOK,
	})
	utils.RenderTemplate(c, http.StatusOK, "players.html", gin.H{
		"success":   fmt.Sprintf("Jugador %s comprado con éxito", player.Name),
		"Players":   playersFormatted,
		"Positions": models.GetAvailablePositions(),
	})
}
