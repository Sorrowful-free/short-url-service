package handler

import (
	"io"
	"net/http"
	"net/url"

	"github.com/labstack/echo/v4"
)

func RegisterMakeShortHandler(h *Handlers) {
	h.internalEcho.POST(MakeShortPath, func(c echo.Context) error {
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

		c.Response().Header().Set(HeaderContentType, HeaderContentTypeText)
		c.Response().WriteHeader(http.StatusCreated)

		return c.String(http.StatusCreated, shortURL)
	})
}
