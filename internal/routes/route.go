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

	// Rutas que necesitan estar autenticado
	authOnly := r.Group("/")
	authOnly.Use(middlewares.AuthRequired())
	{
		authOnly.GET("/profile", controllers.Profile)
		authOnly.GET("/teams/create", controllers.CreateTeamForm)
		authOnly.POST("/teams/create", controllers.CreateTeam)
	}

	// Rutas solo ADMIN
	adminOnly := r.Group("/")
	adminOnly.Use(middlewares.AuthRequired(), middlewares.AdminOnly())
	{
		adminOnly.GET("/tournaments/create", controllers.CreateTournamentForm)
		adminOnly.POST("/tournaments/create", controllers.CreateTournament)
		adminOnly.GET("/players/create", controllers.CreatePlayerForm)
		adminOnly.POST("/players/create", controllers.CreatePlayer)
	}

	// Listado de torneos
	r.GET("/tournaments", controllers.ListTournaments)

	// Listado de equipos
	r.GET("/teams", controllers.ListTeams)

	// Listado de jugadores
	r.GET("/players", controllers.ListPlayers)

}
