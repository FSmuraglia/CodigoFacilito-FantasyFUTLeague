package utils

import (
	"fmt"
	"math/rand"
	"os"

	log "github.com/FSmuraglia/CodigoFacilito-FantasyFUTLeague/internal/logger"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	UserID uint   `json:"user_id"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

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

func SimulateMatch(teamARating, teamBRating float64) string {
	probTeamA := teamARating / (teamARating + teamBRating)
	random := rand.Float64()

	if random < probTeamA {
		return "A"
	}
	return "B"
}

func FormatNumber(n int64) string {
	s := fmt.Sprintf("%d", n)
	formatted := ""
	count := 0
	for i := len(s) - 1; i >= 0; i-- {
		formatted = string(s[i]) + formatted
		count++
		if count%3 == 0 && i != 0 {
			formatted = "." + formatted
		}
	}
	return formatted
}

func GetUserIDFromCookie(c *gin.Context) (uint, bool) {
	tokenString, err := c.Cookie("jwt")
	if err != nil {
		return 0, false
	}

	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil || !token.Valid {
		return 0, false
	}

	if claims, ok := token.Claims.(*Claims); ok {
		return claims.UserID, true
	}

	return 0, false
}

func GetUserRoleFromCookie(c *gin.Context) (string, bool) {
	tokenString, err := c.Cookie("jwt")
	if err != nil {
		return "", false
	}

	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil || !token.Valid {
		return "", false
	}

	if claims, ok := token.Claims.(*Claims); ok {
		return claims.Role, true
	}

	return "", false
}

func GetFormationRequirements(formation string) map[string]int {
	switch formation {
	case "433":
		return map[string]int{
			"Arquero":                 1,
			"Lateral Derecho":         1,
			"Defensor Central":        2,
			"Lateral Izquierdo":       1,
			"Mediocampista Defensivo": 1,
			"Mediocampista Central":   2,
			"Extremo Izquierdo":       1,
			"Extremo Derecho":         1,
			"Delantero Centro":        1,
		}
	case "442":
		return map[string]int{
			"Arquero":                     1,
			"Lateral Derecho":             1,
			"Defensor Central":            2,
			"Lateral Izquierdo":           1,
			"Mediocampista Central":       2,
			"Mediocampista Por Derecha":   1,
			"Mediocampista Por Izquierda": 1,
			"Delantero Centro":            2,
		}
	case "4231":
		return map[string]int{
			"Arquero":                 1,
			"Lateral Derecho":         1,
			"Defensor Central":        2,
			"Lateral Izquierdo":       1,
			"Mediocampista Defensivo": 2,
			"Mediocampista Ofensivo":  1,
			"Extremo Izquierdo":       1,
			"Extremo Derecho":         1,
			"Delantero Centro":        1,
		}
	default:
		return map[string]int{}
	}
}
