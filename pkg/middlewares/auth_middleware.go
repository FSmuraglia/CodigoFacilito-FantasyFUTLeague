package middlewares

import (
	"net/http"
	"os"

	log "github.com/FSmuraglia/CodigoFacilito-FantasyFUTLeague/internal/logger"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		path := c.FullPath()
		tokenString, err := c.Cookie("jwt")
		if err != nil || tokenString == "" {
			log.LogWarn("⚠️ Intento de acceso sin token", map[string]interface{}{
				"path":   path,
				"status": http.StatusUnauthorized,
				"error":  err.Error(),
			})
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Token Requerido",
			})
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET")), nil
		})
		if err != nil || !token.Valid {
			log.LogError("❌ Token inválido", map[string]interface{}{
				"path":   path,
				"status": http.StatusUnauthorized,
				"error":  err.Error(),
			})
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Token Inválido",
			})
			return
		}

		claims := token.Claims.(jwt.MapClaims)
		userID := claims["user_id"]
		role := claims["role"]

		log.LogInfo("✅ Autenticación exitosa", map[string]interface{}{
			"user_id": userID,
			"role":    role,
			"path":    path,
		})

		c.Set("user_id", userID)
		c.Set("role", role)
		c.Next()
	}
}
