package handler

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func RegisterMakeOriginalHandler(h *Handlers) {
	h.internalEcho.GET("/:id", func(c echo.Context) error {
		shortUID := c.Param("id")
		originalURL, err := h.internalURLService.TryMakeOriginal(shortUID)
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}
		c.Response().Header().Set("Content-Type", "text/plain; charset=utf-8")
		c.Response().Header().Set("Location", originalURL)

		fmt.Printf("process request for short UID:%s, with result:%s\n", shortUID, originalURL)
		return c.Redirect(http.StatusTemporaryRedirect, originalURL)
	})
}
