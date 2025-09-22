package handler

import (
	"net/http"

	"github.com/Sorrowful-free/short-url-service/internal/model"
	"github.com/labstack/echo/v4"
)

func (h *Handlers) RegisterGetUserUrlsHandler() {
	h.internalEcho.GET(GetUserPath, func(c echo.Context) error {

		if !h.HasValidUserID(c) {
			return c.String(http.StatusUnauthorized, "unauthorized")
		}

		userID := h.GetUserID(c)

		urls, err := h.internalURLService.GetUserUrls(c.Request().Context(), userID)
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}

		if len(urls) == 0 {
			return c.String(http.StatusNoContent, "no content")
		}

		var getUserUrlsResponse model.GetUserUrlsResponse = urls
		return c.JSON(http.StatusOK, getUserUrlsResponse)
	})
}
