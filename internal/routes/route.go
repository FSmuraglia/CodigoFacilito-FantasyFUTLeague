package routes

import (
	"github.com/FSmuraglia/CodigoFacilito-FantasyFUTLeague/internal/controllers"
	"github.com/FSmuraglia/CodigoFacilito-FantasyFUTLeague/pkg/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	// Página principal
	r.GET("/", controllers.ShowIndex)

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
		authOnly.GET("/profile", controllers.GetProfile)
		authOnly.GET("/teams", controllers.ListTeams)
		authOnly.GET("/teams/create", controllers.CreateTeamForm)
		authOnly.POST("/teams/create", controllers.CreateTeam)
		authOnly.GET("/tournaments/:id", controllers.GetTournamentDetail)
		authOnly.POST("/tournaments/:id/join", controllers.JoinTournament)
		authOnly.GET("/tournaments", controllers.ListTournaments)
		authOnly.GET("/players", controllers.ListPlayers)
		authOnly.POST("/players/:id/buy", controllers.BuyPlayer)
		authOnly.GET("/matches", controllers.ListMatches)
		authOnly.GET("/matches/:id", controllers.GetMatchDetail)
	}

	// Rutas solo ADMIN
	adminOnly := r.Group("/")
	adminOnly.Use(middlewares.AuthRequired(), middlewares.AdminOnly())
	{
		adminOnly.GET("/tournaments/create", controllers.CreateTournamentForm)
		adminOnly.POST("/tournaments/create", controllers.CreateTournament)
		adminOnly.POST("/tournaments/:id/finish", controllers.FinishTournament)
		adminOnly.GET("/players/create", controllers.CreatePlayerForm)
		adminOnly.POST("/players/create", controllers.CreatePlayer)
		adminOnly.GET("/matches/create", controllers.CreateMatchForm)
		adminOnly.POST("/matches/create", controllers.CreateMatch)
		adminOnly.GET("/tournaments/:id/teams", controllers.GetTeamsByTournament)
		adminOnly.POST("/matches/:id/simulate", controllers.SimulateMatchController)
	}
}
