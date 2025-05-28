package logger

import (
	"go-tutorial/internal/config"
	"go.uber.org/zap"
)

func New(cfg *config.Config) *zap.Logger {
	var logger *zap.Logger

	if cfg.Server.IsProduction() {
		logger, _ = zap.NewProduction()
	} else {
		logger, _ = zap.NewDevelopment()
	}

	return logger
}
