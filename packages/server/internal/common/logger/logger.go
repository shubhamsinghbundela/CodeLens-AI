package logger

import (
	envConfig "my-codelens-app/internal/config"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.Logger

// Initialize sets up the global logger with the provided configuration
// Example usage:
//
//	cfg := &logger.Config{
//	    Environment: "development",
//	    Level: "info",
//	}
//	logger.Initialize(cfg)
//	logger.Info("Application started")
func Initialize(config envConfig.Config) error {
	var zapConfig zap.Config

	if config.AppEnv != envConfig.Development {
		zapConfig = zap.NewProductionConfig()
		zapConfig.EncoderConfig.TimeKey = "timestamp"
		zapConfig.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	} else {
		zapConfig = zap.NewDevelopmentConfig()
		zapConfig.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	}

	// Set log level
	level, err := zapcore.ParseLevel(config.LogLevel)
	if err != nil {
		level = zapcore.InfoLevel
	}
	zapConfig.Level = zap.NewAtomicLevelAt(level)
	zapConfig.DisableCaller = true
	zapConfig.DisableStacktrace = true

	logger, err := zapConfig.Build()
	if err != nil {
		return err
	}

	Logger = logger
	return nil
}

// Sync flushes any buffered log entries
func Sync() {
	if Logger != nil {
		if err := Logger.Sync(); err != nil {
			Logger.Error("Failed to sync logger", zap.Error(err))
		}
	}
}

// Convenience functions for logging
func Debug(msg string, fields ...zap.Field) {
	if Logger != nil {
		Logger.Debug(msg, fields...)
	}
}

func Info(msg string, fields ...zap.Field) {
	if Logger != nil {
		Logger.Info(msg, fields...)
	}
}

func Warn(msg string, fields ...zap.Field) {
	if Logger != nil {
		Logger.Warn(msg, fields...)
	}
}

func Error(msg string, fields ...zap.Field) {
	if Logger != nil {
		Logger.Error(msg, fields...)
	}
}

func Fatal(msg string, fields ...zap.Field) {
	if Logger != nil {
		Logger.Fatal(msg, fields...)
	}
}

func Panic(msg string, fields ...zap.Field) {
	if Logger != nil {
		Logger.Panic(msg, fields...)
	}
}

// WithFields returns a logger with predefined fields
func WithFields(fields ...zap.Field) *zap.Logger {
	if Logger != nil {
		return Logger.With(fields...)
	}
	return nil
}
