package middlewares

import (
	"time"

	"github.com/Sorrowful-free/short-url-service/internal/logger"
	"github.com/labstack/echo/v4"
)

func LoggerAsMiddleware(l *logger.Logger) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {

			begin := time.Now()
			err := next(c)
			duration := time.Since(begin)

			if err != nil {
				l.Error("error processing request: ",
					"url", c.Request().URL.Path,
					"method", c.Request().Method,
					"headers", c.Request().Header,
					"error", err,
					"duration", duration)
			} else {
				l.Info("request processed: ",
					"url", c.Request().URL.Path,
					"method", c.Request().Method,
					"headers", c.Request().Header,
					"duration", duration,
				)
				l.Info("response processed: ",
					"code", c.Response().Status,
					"headers", c.Response().Header(),
					"size", c.Response().Size)
			}
			return err
		}
	}
}
