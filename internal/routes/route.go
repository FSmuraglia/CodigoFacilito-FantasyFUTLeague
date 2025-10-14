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

	//Logout
	r.GET("/logout", controllers.LogoutUser)

	// Perfil (requiere autenticación)
	protected := r.Group("/")
	protected.Use(middlewares.AuthRequired())
	{
		protected.GET("/profile", controllers.Profile)
		protected.GET("/teams/create", controllers.CreateTeamForm)
		protected.POST("/teams/create", controllers.CreateTeam)
	}

	adminTournament := r.Group("/tournaments")
	adminTournament.Use(middlewares.AuthRequired(), middlewares.AdminOnly())
	{
		adminTournament.GET("/create", controllers.CreateTournamentForm)
		adminTournament.POST("/create", controllers.CreateTournament)
	}

	// Listado de torneos
	r.GET("/tournaments", controllers.ListTournaments)

}
