package logger

import (
	"go.uber.org/zap"
)

type Logger interface {
	Info(message string, args ...interface{})
	Error(message string, args ...interface{})
	Debug(message string, args ...interface{})
}

type ZapLogger struct {
	sugaredLogger *zap.SugaredLogger
}

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
