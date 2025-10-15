package controllers

import (
	"net/http"

	database "github.com/FSmuraglia/CodigoFacilito-FantasyFUTLeague/config"
	log "github.com/FSmuraglia/CodigoFacilito-FantasyFUTLeague/internal/logger"
	"github.com/FSmuraglia/CodigoFacilito-FantasyFUTLeague/internal/models"
	"github.com/FSmuraglia/CodigoFacilito-FantasyFUTLeague/pkg/utils"
	"github.com/gin-gonic/gin"
)

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
