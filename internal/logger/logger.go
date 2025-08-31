package logger

import (
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

func (l *Logger) AsMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			err := next(c)
			if err != nil {
				l.sugaredLogger.Error("request", "method", c.Request().Method, "path", c.Request().URL.Path, "error", err)
			}
			l.sugaredLogger.Info("request", "method", c.Request().Method, "path", c.Request().URL.Path)
			return err
		}
	}
}
