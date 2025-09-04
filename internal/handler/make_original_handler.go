package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func RegisterMakeOriginalHandler(h *Handlers) {
	h.internalEcho.GET(MakeOriginalPath, func(c echo.Context) error {
		shortUID := c.Param(OriginalPathParam)
		originalURL, err := h.internalURLService.TryMakeOriginal(shortUID)
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}
		c.Response().Header().Set(HeaderContentType, HeaderContentTypeText)
		c.Response().Header().Set(HeaderLocation, originalURL)

		return c.Redirect(http.StatusTemporaryRedirect, originalURL)
	})
}
