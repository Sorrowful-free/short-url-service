package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *Handlers) RegisterGetStatHandler() {
	h.internalEcho.GET(GetUserURLsPath, func(c echo.Context) error {
		stats, err := h.internalStatService.GetStats(c.Request().Context())
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, stats)
	})
}
