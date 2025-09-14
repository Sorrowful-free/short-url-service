package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *Handlers) RegisterPingDBHandler() {
	h.internalEcho.GET("/ping", func(c echo.Context) error {
		err := h.internalDBService.Ping()
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}
		return c.String(http.StatusOK, "pong")
	})
}
