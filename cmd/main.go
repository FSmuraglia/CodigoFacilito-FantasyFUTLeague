package main

import (
	"math/rand"
	"text/template"
	"time"

	database "github.com/FSmuraglia/CodigoFacilito-FantasyFUTLeague/config"
	"github.com/FSmuraglia/CodigoFacilito-FantasyFUTLeague/internal/controllers"
	log "github.com/FSmuraglia/CodigoFacilito-FantasyFUTLeague/internal/logger"
	"github.com/FSmuraglia/CodigoFacilito-FantasyFUTLeague/internal/repositories"
	"github.com/FSmuraglia/CodigoFacilito-FantasyFUTLeague/internal/routes"
	"github.com/FSmuraglia/CodigoFacilito-FantasyFUTLeague/internal/services"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	rand.Seed(time.Now().UnixNano())

	err := godotenv.Load()

	if err != nil {
		log.LogWarn("⚠️ No se encontró archivo .env, usando variables del sistema", nil)
	}

	database.Connect()

	// Inicialización de repositorios y services
	tournamentsRepository := repositories.NewTournamentRepository()
	tournamentService := services.NewTournamentService(tournamentsRepository)
	controllers.InitTournamentController(tournamentService)

	teamRepository := repositories.NewTeamRepository()
	teamService := services.NewTeamService(teamRepository)
	controllers.InitTeamController(teamService)

	playerRepository := repositories.NewPlayerRepository()
	playerService := services.NewPlayerService(playerRepository)
	controllers.InitPlayerController(playerService)

	matchRepository := repositories.NewMatchRepository()
	matchService := services.NewMatchService(matchRepository)
	controllers.InitMatchController(matchService)

	userRepository := repositories.NewUserRepository()
	userService := services.NewUserService(userRepository)
	controllers.InitProfileController(userService)

	r := gin.Default()

	r.SetFuncMap(template.FuncMap{
		"add": func(a, b int) int { return a + b },
	})

	r.Static("/static", "./static")

	r.LoadHTMLGlob("templates/*")

	routes.RegisterRoutes(r)

	r.Run(":8080")
}
