package logger

import (
	"fmt"

	"go.uber.org/zap"
)

func NewLogger() (*zap.SugaredLogger, error) {
	cfg := zap.NewDevelopmentConfig()

	logger, err := cfg.Build()
	if err != nil {
		return nil, fmt.Errorf("could not build logger: %w", err)
	}

	zap.ReplaceGlobals(logger)
logger.Sugar().Sync()
	return logger.Sugar(), nil
}
