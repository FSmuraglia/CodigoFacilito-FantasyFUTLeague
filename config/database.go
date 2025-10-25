package config

import (
	"fmt"
	"os"
	"time"

	log "github.com/FSmuraglia/CodigoFacilito-FantasyFUTLeague/internal/logger"
	"github.com/FSmuraglia/CodigoFacilito-FantasyFUTLeague/internal/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	var err error

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	log.LogInfo(fmt.Sprintf("Intentando conectar con DSN: %s", dsn), nil)

	maxRetries := 15
	for i := 1; i <= maxRetries; i++ {
		DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err == nil {
			log.LogInfo(fmt.Sprintf("✅ Conexión establecida en el intento #%d", i), nil)
			break
		}

		log.LogWarn(fmt.Sprintf("⏳ Intento %d de conexión fallido, reintentando...", i),
			map[string]interface{}{"error": err.Error()})
		time.Sleep(5 * time.Second)
	}

	if err != nil {
		log.LogError("❌ No se pudo conectar a la base de datos después de varios intentos", map[string]interface{}{
			"error": err.Error(),
		})
		panic("No se pudo conectar a la base de datos")
	}

	sqlDB, err := DB.DB()
	if err != nil {
		log.LogError("❌ No se pudo obtener instancia SQL del driver", map[string]interface{}{
			"error": err.Error(),
		})
		panic("No se pudo obtener instancia SQL del driver")
	}

	// Configuración de pool
	sqlDB.SetMaxOpenConns(10)
	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetConnMaxLifetime(time.Hour)

	// Migraciones (solo si DB != nil)
	if DB != nil {
		log.LogInfo("📦 Ejecutando migraciones automáticas...", nil)
		err = DB.AutoMigrate(
			&models.User{},
			&models.Team{},
			&models.Player{},
			&models.Tournament{},
			&models.TournamentTeam{},
			&models.Match{},
		)
		if err != nil {
			log.LogError("❌ Error durante las migraciones automáticas", map[string]interface{}{
				"error": err.Error(),
			})
			panic("Error durante las migraciones automáticas")
		}

		log.LogInfo("✅ Migraciones completadas correctamente", nil)
	}

}
