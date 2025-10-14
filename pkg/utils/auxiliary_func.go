package utils

import (
	"os"

	log "github.com/FSmuraglia/CodigoFacilito-FantasyFUTLeague/internal/logger"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
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

func GetUserIDFromJWT(c *gin.Context) (uint, bool) {
	cookie, err := c.Cookie("jwt")
	if err != nil {
		return 0, false
	}

	token, err := jwt.Parse(cookie, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil || !token.Valid {
		return 0, false
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0, false
	}

	idFloat, ok := claims["user_id"].(float64)
	if !ok {
		return 0, false
	}

	return uint(idFloat), true

}
