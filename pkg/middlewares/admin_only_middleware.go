package middlewares

import (
	"net/http"

	log "github.com/FSmuraglia/CodigoFacilito-FantasyFUTLeague/internal/logger"
	"github.com/gin-gonic/gin"
)

func AdminOnly() gin.HandlerFunc {
	return func(c *gin.Context) {
		path := c.FullPath()
		role, _ := c.Get("role")
		if role != "ADMIN" {
			log.LogWarn("⚠️ Acceso denegado a ruta admin", map[string]interface{}{
				"path":   path,
				"role":   role,
				"status": http.StatusForbidden,
			})
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"error": "Acceso Denegado",
			})
			return
		}
		log.LogInfo("✅ Acceso autorizado a ruta admin", map[string]interface{}{
			"path": path,
		})
		c.Next()
	}
}
