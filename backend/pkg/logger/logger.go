package logger

import (
	"github.com/phuslu/log"
)

// Init configures the default logger for the application.
func Init() {
	log.DefaultLogger = log.Logger{
		Level:      log.InfoLevel,
		Caller:     1,
		TimeFormat: "2006-01-02 15:04:05",
		Writer:     &log.ConsoleWriter{},
	}
}
