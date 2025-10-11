package routes

import (
	"github.com/FSmuraglia/CodigoFacilito-FantasyFUTLeague/internal/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	// Página principal
	r.GET("/", controllers.Index)

	// Register
	r.GET("/register", controllers.RegisterForm)
	r.POST("/register", controllers.RegisterUser)

	// Login
	r.GET("/login", controllers.LoginForm)
	r.POST("/login", controllers.LoginUser)

	// Perfil (requiere autenticación)
	r.GET("/profile", controllers.Profile)
}
