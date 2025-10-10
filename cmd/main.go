package main

import (
	database "github.com/FSmuraglia/CodigoFacilito-FantasyFUTLeague/config"
	"github.com/FSmuraglia/CodigoFacilito-FantasyFUTLeague/internal/routes"
	"github.com/gin-gonic/gin"
)

func main() {

	database.Connect()

	r := gin.Default()

	r.LoadHTMLGlob("templates/*")

	routes.RegisterRoutes(r)

	r.Run(":8080")
}
