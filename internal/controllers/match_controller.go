package controllers

import (
	"net/http"

	log "github.com/FSmuraglia/CodigoFacilito-FantasyFUTLeague/internal/logger"
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

	utils.RenderTemplate(c, http.StatusOK, "matches.html", gin.H{
		"matches": matches,
		"Sort":    sort,
		"Status":  status,
	})

}
