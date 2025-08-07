package services

import (
	"github.com/go-logr/logr"
	"github.com/go-logr/zapr"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewLogger(logFormat string) logr.Logger {
	if logFormat == "json" {
		return NewJSONLogger()
	} else {
		return NewTextLogger()
	}
}

func NewJSONLogger() logr.Logger {
	cfg := zap.NewProductionConfig() // JSON output

	zapLogger, _ := cfg.Build()
	return zapr.NewLogger(zapLogger)
}

func NewTextLogger() logr.Logger {
	cfg := zap.NewDevelopmentConfig()                                // Human-readable output
	cfg.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder // colored levels

	zapLogger, _ := cfg.Build()
	return zapr.NewLogger(zapLogger)
}
