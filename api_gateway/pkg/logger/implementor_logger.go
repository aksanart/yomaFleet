package logger

import (
	"log"

	"github.com/aksan/weplus/apigw/pkg/config"
)

var (
	logger *Logger
)

func LoadLogger() {
	appLogger, err := NewLogger(LoggerConfig{
		Level:           config.OsGetString("LOG_LEVEL", "debug"),
		LogDirectory:    config.OsGetString("LOG_DIRECTORY", ""),
		AppName:         config.OsGetString("APP_NAME", ""),
		SamplingEnabled: false,
	})
	if err != nil {
		log.Fatalf("cannot initiate logger, with error: %v", err)
	}
	logger = appLogger
}

func GetLogger() *Logger {
	return logger
}
