package logger

import (
	"go.uber.org/zap"
)

type Logger struct {
	sugaredLogger *zap.SugaredLogger
}

func NewLogger() (*Logger, error) {
	logger, err := zap.NewDevelopment()
	if err != nil {
		return nil, err
	}
	return &Logger{
		sugaredLogger: logger.Sugar(),
	}, nil
}

func (l *Logger) Info(message string, args ...interface{}) {
	l.sugaredLogger.Infow(message, args...)
}

func (l *Logger) Error(message string, args ...interface{}) {
	l.sugaredLogger.Errorw(message, args...)
}

func (l *Logger) Debug(message string, args ...interface{}) {
	l.sugaredLogger.Debugw(message, args...)
}
