package routes

import (
	"github.com/FSmuraglia/CodigoFacilito-FantasyFUTLeague/internal/controllers"
	"github.com/FSmuraglia/CodigoFacilito-FantasyFUTLeague/pkg/middlewares"
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
	protected := r.Group("/")
	protected.Use(middlewares.AuthRequired())
	{
		protected.GET("/profile", controllers.Profile)
	}

	// Rutas Solo ADMIN
	/*
		admin := r.Group("/admin")
		admin.Use(middlewares.AuthRequired(), middlewares.AdminOnly())
		{
			admin.POST("/tournaments", controllers.CreateTournament)
		}
	*/
}
