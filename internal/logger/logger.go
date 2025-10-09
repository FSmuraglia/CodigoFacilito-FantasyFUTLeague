package logger

import (
	"os"

	log "github.com/sirupsen/logrus"
)

var Logger *log.Logger

func init() {
	Logger = log.New()

	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666) // Permisos Usuario Grupo Otros, Lectura + Escritura para todos
	if err != nil {
		log.Fatalf("No se pudo abrir el archivo de logs %v", err)
	}
	Logger.Out = file

	Logger.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})

	Logger.SetLevel(log.DebugLevel)
}

func LogInfo(msg string, fields log.Fields) {
	Logger.WithFields(fields).Info(msg)
}

func LogError(msg string, fields log.Fields) {
	Logger.WithFields(fields).Error(msg)
}

func LogDebug(msg string, fields log.Fields) {
	Logger.WithFields(fields).Debug(msg)
}

func LogWarn(msg string, fields log.Fields) {
	Logger.WithFields(fields).Warn(msg)
}

func LogFatal(msg string, fields log.Fields) {
	Logger.WithFields(fields).Fatal(msg)
}
