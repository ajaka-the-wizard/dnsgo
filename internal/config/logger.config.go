package config

import (
	"log"

	"go.uber.org/zap"
)

func InitializeLogger(env *Env) *zap.Logger {
	var logger *zap.Logger
	var err error
	if env.PRODUCTION {
		logger, err = zap.NewProduction()
	} else {
		logger, err = zap.NewDevelopment()
	}
	if err != nil {
		log.Fatalf("Couldn't init logger :%v", err)
	}
	zap.ReplaceGlobals(logger)
	return logger
}
