package handler

import (
	"encoding/json"
	"net/http"

	"github.com/Sorrowful-free/short-url-service/internal/middlewares"
	"github.com/Sorrowful-free/short-url-service/internal/model"
	"github.com/labstack/echo/v4"
)

func (h *Handlers) RegisterDeleteUserURLsHandler() {
	h.internalEcho.DELETE(DeleteUserURLsPath, func(c echo.Context) error {
		userID := middlewares.TryGetUserID(c)

		var deleteShortURLRequest model.DeleteShortURLRequest
		dec := json.NewDecoder(c.Request().Body)
		err := dec.Decode(&deleteShortURLRequest)
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}
		err = h.internalURLService.DeleteShortURLs(c.Request().Context(), userID, deleteShortURLRequest)
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}
		return c.String(http.StatusAccepted, "accepted")
	})
}
