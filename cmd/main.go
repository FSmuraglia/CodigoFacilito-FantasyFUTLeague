package main

import (
	"github.com/FSmuraglia/CodigoFacilito-FantasyFUTLeague/internal/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.LoadHTMLGlob("templates/*")

	routes.RegisterRoutes(r)

	r.Run(":8080")
}
