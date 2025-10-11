package config

import (
	"fmt"
	"os"

	log "github.com/FSmuraglia/CodigoFacilito-FantasyFUTLeague/internal/logger"
	"github.com/FSmuraglia/CodigoFacilito-FantasyFUTLeague/internal/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func GetConnection() (*gorm.DB, error) {
	err := godotenv.Load()

	if err != nil {
		log.LogWarn("⚠️ No se encontró archivo .env, usando variables del sistema", nil)
	}

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	if dsn == "" {
		log.LogError("❌ No se encontraron una o varias variables en el entorno", nil)
	}

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	return db, nil
}

func Connect() {
	db, err := GetConnection()

	if err != nil {
		log.LogError("❌ Error al conectar a MySQL", nil)
	}

	log.LogInfo("✅ Conexión a base de datos exitosa", nil)

	if err := db.AutoMigrate(
		&models.User{},
		&models.Team{},
		&models.Tournament{},
		&models.Player{},
		&models.Match{},
		&models.TournamentTeam{}); err != nil {
		log.LogError("❌ Error al migrar modelos", map[string]interface{}{
			"error": err.Error(),
		})
	} else {
		log.LogInfo("✅ Migraciones completadas", nil)
	}
}
