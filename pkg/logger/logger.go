package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewLogger() *zap.Logger {
	var cfg zap.Config = zap.NewProductionConfig()
	cfg.EncoderConfig.EncodeTime = zapcore.RFC3339TimeEncoder
	logger, err := cfg.Build()

	if err != nil {
		panic(err)
	}

	return logger
}
