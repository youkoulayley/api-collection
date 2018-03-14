package bootstrap

import (
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/youkoulayley/api-collection/models"
)

// InitLogs set options for logs
func InitLogs(c *models.Conf) {
	l := c.Log

	// Log as JSON instead of the default ASCII formatter
	if l.Formatter == "json" {
		log.SetFormatter(&log.JSONFormatter{})
		log.Info("Logs - Set to JSON Format")
	}

	// Output to stdout instead of the default stderr
	switch l.Output {
	case "stdout":
		log.SetOutput(os.Stdout)
	case "sterr":
		log.SetOutput(os.Stderr)
	}

	log.Info("Logs - Output set to ", l.Output)

	// Only log the warning severity or above.
	switch l.Level {
	case "debug":
		log.SetLevel(log.DebugLevel)
	case "info":
		log.SetLevel(log.InfoLevel)
	case "warn":
		log.SetLevel(log.InfoLevel)
	case "crit":
		log.SetLevel(log.InfoLevel)
	}
	log.Info("Logs - Level set to ", l.Level)
}
