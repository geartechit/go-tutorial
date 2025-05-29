package logger

import (
	"go-tutorial/internal/config"
	"go.uber.org/zap"
)

type ZapLogger struct {
	z *zap.Logger
}

func NewZapLogger(cfg *config.Config) *ZapLogger {
	var z *zap.Logger

	if cfg.Server.IsProduction() {
		z, _ = zap.NewProduction()
	} else {
		z, _ = zap.NewDevelopment()
	}
	return &ZapLogger{z: z}
}

func (l *ZapLogger) Error(msg string, fields ...Field) {
	l.z.Error(msg, mapZapField(fields...)...)
}
func (l *ZapLogger) Info(msg string, fields ...Field) {
	l.z.Info(msg, mapZapField(fields...)...)
}
func (l *ZapLogger) Warn(msg string, fields ...Field) {
	l.z.Warn(msg, mapZapField(fields...)...)
}

func mapZapField(fields ...Field) []zap.Field {
	field := make([]zap.Field, len(fields))
	for i, f := range fields {
		field[i] = zap.Any(f.Key, f.Value)
	}
	return field
}
