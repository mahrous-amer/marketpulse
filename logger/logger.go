package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

// Logger instance
var log = logrus.New()

// SetLogLevel sets the logging level
func SetLogLevel(level string) {
	switch level {
	case "debug":
		log.SetLevel(logrus.DebugLevel)
	case "info":
		log.SetLevel(logrus.InfoLevel)
	case "warn":
		log.SetLevel(logrus.WarnLevel)
	case "error":
		log.SetLevel(logrus.ErrorLevel)
	default:
		log.SetLevel(logrus.InfoLevel)
	}

	log.SetOutput(os.Stdout)
}

// Expose functions per level

func Info(msg string) {
	log.Info(msg)
}

func Error(msg string) {
	log.Error(msg)
}

func Warn(msg string) {
	log.Warn(msg)
}
