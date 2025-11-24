package logger

import (
	"go.uber.org/zap"
)

// Logger defines the interface for application logging.
// It provides methods for logging at different severity levels.
type Logger interface {
	// Info logs an informational message with optional key-value pairs.
	Info(message string, args ...interface{})

	// Error logs an error message with optional key-value pairs.
	Error(message string, args ...interface{})

	// Debug logs a debug message with optional key-value pairs.
	Debug(message string, args ...interface{})
}

// ZapLogger is a zap-based implementation of the Logger interface.
// It uses zap's SugaredLogger for structured logging.
type ZapLogger struct {
	sugaredLogger *zap.SugaredLogger
}

// NewZapLogger creates a new zap-based logger instance.
// It initializes a development logger with default settings.
// Returns a Logger implementation and an error if initialization fails.
func NewZapLogger() (Logger, error) {
	logger, err := zap.NewDevelopment()
	if err != nil {
		return nil, err
	}
	return &ZapLogger{
		sugaredLogger: logger.Sugar(),
	}, nil
}

func (l *ZapLogger) Info(message string, args ...interface{}) {
	l.sugaredLogger.Infow(message, args...)
}

func (l *ZapLogger) Error(message string, args ...interface{}) {
	l.sugaredLogger.Errorw(message, args...)
}

func (l *ZapLogger) Debug(message string, args ...interface{}) {
	l.sugaredLogger.Debugw(message, args...)
}
