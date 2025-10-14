package controllers

import (
	"net/http"
	"path/filepath"
	"strings"
	"time"

	database "github.com/FSmuraglia/CodigoFacilito-FantasyFUTLeague/config"
	log "github.com/FSmuraglia/CodigoFacilito-FantasyFUTLeague/internal/logger"
	"github.com/FSmuraglia/CodigoFacilito-FantasyFUTLeague/internal/models"
	"github.com/FSmuraglia/CodigoFacilito-FantasyFUTLeague/internal/services"
	"github.com/FSmuraglia/CodigoFacilito-FantasyFUTLeague/pkg/utils"
	"github.com/gin-gonic/gin"
)

var teamService *services.TeamService

func InitTeamController(s *services.TeamService) {
	teamService = s
}

func CreateTeamForm(c *gin.Context) {
	log.LogInfo("📝 Acceso a formulario de registro de equipo", nil)
	utils.RenderTemplate(c, http.StatusOK, "create_team.html", gin.H{
		"Formations": models.GetAvailableFormations(),
	})
}

func CreateTeam(c *gin.Context) {
	userID, exists := utils.GetUserIDFromJWT(c)
	if !exists {
		log.LogWarn("⚠️ Usuario no autenticado intentando crear equipo", nil)
		c.Redirect(http.StatusSeeOther, "/login")
		return
	}

	name := c.PostForm("name")
	formation := c.PostForm("formation")

	// Procesar imagen del escudo
	file, err := c.FormFile("badge")
	if err != nil {
		log.LogWarn("⚠️ Escudo no subido", map[string]interface{}{
			"status":  http.StatusBadRequest,
			"user_id": userID,
		})
		utils.RenderTemplate(c, http.StatusBadRequest, "create_team.html", gin.H{
			"error":      "Debe subir una imagen de escudo",
			"Formations": models.GetAvailableFormations(),
		})
		return
	}

	// Guardar imagen en carpeta /static/uploads
	filename := time.Now().Format("20060102150405") + "_" + filepath.Base(file.Filename)
	savePath := filepath.Join("static", "uploads", filename)
	if err := c.SaveUploadedFile(file, savePath); err != nil {
		log.LogError("❌ Error al guardar la imagen de escudo", map[string]interface{}{
			"error":   err.Error(),
			"status":  http.StatusInternalServerError,
			"user_id": userID,
		})
		utils.RenderTemplate(c, http.StatusInternalServerError, "create_team.html", gin.H{
			"error":      "Error al guardar imagen",
			"Formations": models.GetAvailableFormations(),
		})
		return
	}

	// Crear el equipo
	team := models.Team{
		Name:      name,
		Formation: models.Formation(formation),
		BadgeUrl:  "/static/uploads/" + filename,
		UserID:    userID,
	}

	if err := database.DB.Create(&team).Error; err != nil {
		log.LogError("❌ Error al crear el equipo en la DB", map[string]interface{}{
			"error":   err.Error(),
			"status":  http.StatusInternalServerError,
			"user_id": userID,
		})
		utils.RenderTemplate(c, http.StatusInternalServerError, "create_team.html", gin.H{
			"error":      "Error al crear equipo",
			"Formations": models.GetAvailableFormations(),
		})
		return
	}

	log.LogInfo("✅ Equipo creado correctamente", map[string]interface{}{
		"user_id": userID,
		"team":    team.Name,
	})

	c.Redirect(http.StatusSeeOther, "/")

}

func ListTeams(c *gin.Context) {
	nameFilter := strings.TrimSpace(c.Query("name"))
	formationFilter := c.Query("formation")

	teams, err := teamService.ListTeams(nameFilter, formationFilter)
	if err != nil {
		log.LogError("❌ Error al obtener los equipos de la DB", map[string]interface{}{
			"status": http.StatusInternalServerError,
			"error":  err.Error(),
		})
		c.HTML(http.StatusInternalServerError, "teams.html", gin.H{
			"error": "Error al obtener los equipos de la DB",
		})
		return
	}

	log.LogInfo("✅ Equipos obtenidos correctamente de la DB", map[string]interface{}{
		"count":  len(teams),
		"status": http.StatusOK,
	})
	utils.RenderTemplate(c, http.StatusOK, "teams.html", gin.H{
		"teams":           teams,
		"NameFilter":      nameFilter,
		"FormationFilter": formationFilter,
		"Formations":      models.GetAvailableFormations(),
	})
}
