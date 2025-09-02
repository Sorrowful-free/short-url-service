package logger

import (
	"time"

	"github.com/labstack/echo/v4"
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

func (l *Logger) AsMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {

			begin := time.Now()
			err := next(c)
			duration := time.Since(begin)

			if err != nil {
				l.Error("error processing request: ",
					"url", c.Request().URL.Path,
					"method", c.Request().Method,
					"error", err,
					"duration", duration)
			} else {
				l.Info("request processed: ",
					"url", c.Request().URL.Path,
					"method", c.Request().Method,
					"duration", duration,
				)
				l.Info("response processed: ",
					"code", c.Response().Status,
					"size", c.Response().Size)
			}
			return err
		}
	}
}
