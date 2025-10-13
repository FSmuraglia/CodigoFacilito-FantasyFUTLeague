package controllers

import (
	"net/http"
	"time"

	database "github.com/FSmuraglia/CodigoFacilito-FantasyFUTLeague/config"
	log "github.com/FSmuraglia/CodigoFacilito-FantasyFUTLeague/internal/logger"
	"github.com/FSmuraglia/CodigoFacilito-FantasyFUTLeague/internal/models"
	"github.com/gin-gonic/gin"
)

func CreateTournamentForm(c *gin.Context) {
	log.LogInfo("📝 Acceso a formulario de registro de torneo", nil)
	c.HTML(http.StatusOK, "create_tournament.html", nil)
}

func CreateTournament(c *gin.Context) {
	var input struct {
		Name      string  `form:"name" binding:"required"`
		Prize     float64 `form:"prize" binding:"required"`
		StartDate string  `form:"start_date" binding:"required"`
		EndDate   string  `form:"end_date"`
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
		Name:      input.Name,
		Prize:     input.Prize,
		StartDate: startDate,
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
