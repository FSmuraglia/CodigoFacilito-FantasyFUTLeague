package utils

import (
	log "github.com/FSmuraglia/CodigoFacilito-FantasyFUTLeague/internal/logger"
	"github.com/gin-gonic/gin"
)

// Función para verificar la autenticación del usuario y renderizar la vista
func RenderTemplate(c *gin.Context, status int, templateName string, data gin.H) {
	token, err := c.Cookie("jwt")
	if err != nil {
		log.LogWarn("🔍 Cookie jwt no encontrada, el usuario está deslogeado", nil)
	}

	isLoggedIn := err == nil && token != ""

	// Si no llegan más cosas en el gin.H por parámetro, creamos un map vacío
	if data == nil {
		data = gin.H{}
	}

	// Se incluye siempre el isLoggedIn en el gin.H{}
	data["IsLoggedIn"] = isLoggedIn

	c.HTML(status, templateName, data)
}
