package middlewares

import (
	"net/http"
	"strings"

	"github.com/Sorrowful-free/short-url-service/internal/compression"
	"github.com/Sorrowful-free/short-url-service/internal/consts"
	"github.com/Sorrowful-free/short-url-service/internal/logger"
	"github.com/labstack/echo/v4"
)

func GzipMiddleware(logger *logger.Logger) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {

			contentEncoding := c.Request().Header.Get(consts.HeaderContentEncoding)
			isGzipRequested := strings.Contains(contentEncoding, consts.HeaderEncodingGzip)

			contentType := c.Request().Header.Get(consts.HeaderContentType)

			if isGzipRequested {
				gzr, err := compression.NewGzipRequestReader(c.Request())
				if err != nil {
					return c.String(http.StatusInternalServerError, err.Error())
				}
				defer gzr.Close()
				c.Request().Body = gzr
				logger.Info("gzip request reader created", "contentEncoding", contentEncoding, "isGzipRequested", isGzipRequested, "isSupportedContent", isSupportedContent)
			}

			acceptEncoding := c.Request().Header.Get(consts.HeaderAcceptEncoding)
			isGzipAccepted := strings.Contains(acceptEncoding, consts.HeaderEncodingGzip)
			contentType = c.Response().Header().Get(consts.HeaderContentType)
			isAcceptedContent := acceptContentType(contentType)

			if isGzipAccepted && isAcceptedContent {
				gzw := compression.NewGzipResponseWriter(c.Response())
				defer gzw.Close()
				c.Response().Writer = gzw
				logger.Info("gzip response writer created", "contentType", contentType, "isGzipAccepted", isGzipAccepted, "isAcceptedContent", isAcceptedContent)
			}

			return next(c)
		}
	}
}

func acceptContentType(contentType string) bool {
	return strings.Contains(contentType, consts.HeaderContentTypeHTML) || strings.Contains(contentType, consts.HeaderContentTypeJSON)
}
