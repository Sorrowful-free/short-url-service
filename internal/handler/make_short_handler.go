package handler

import (
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/labstack/echo/v4"
)

func RegisterMakeShortHandler(h *Handlers) {
	h.internalEcho.POST("/", func(c echo.Context) error {
		originalURL, err := io.ReadAll(c.Request().Body)
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}
		shortUID, err := h.internalURLService.TryMakeShort(string(originalURL))
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}

		shortURL, err := url.JoinPath(h.internalBaseURL, shortUID)
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}

		c.Response().Header().Set("Content-Type", "text/plain")
		c.Response().WriteHeader(http.StatusCreated)

		fmt.Printf("process request for original URL:%s, with result:%s\n", originalURL, shortURL)
		return c.String(http.StatusCreated, shortURL)
	})
}
