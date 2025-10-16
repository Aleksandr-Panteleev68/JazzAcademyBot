package utils

import "go.uber.org/zap"

func NewLogger(level string) *zap.Logger {
	config := zap.NewProductionConfig()
	switch level {
	case "debug":
		config.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	case "error":
		config.Level = zap.NewAtomicLevelAt(zap.ErrorLevel)
	default:
		config.Level = zap.NewAtomicLevelAt(zap.InfoLevel)
	}
	logger, _ := config.Build()
	return logger
}
