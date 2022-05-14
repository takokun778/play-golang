package logger

import (
	"os"

	"go.uber.org/zap"
)

var logger *zap.Logger

func init() {
	level := zap.InfoLevel

	if os.Getenv("ENV") == "dev" {
		level = zap.DebugLevel
	}

	config := zap.Config{
		Level:            zap.NewAtomicLevelAt(level),
		Development:      true,
		Encoding:         "console",
		EncoderConfig:    zap.NewDevelopmentEncoderConfig(),
		OutputPaths:      []string{"stderr"},
		ErrorOutputPaths: []string{"stderr"},
	}
	logger, _ = config.Build()
	defer logger.Sync()
}

func Logger() *zap.SugaredLogger {
	return logger.Sugar()
}
