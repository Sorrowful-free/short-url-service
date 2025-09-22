package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"

	"github.com/Sorrowful-free/short-url-service/internal/consts"
	"github.com/Sorrowful-free/short-url-service/internal/model"
	"github.com/Sorrowful-free/short-url-service/internal/service"
	"github.com/labstack/echo/v4"
)

func (h *Handlers) RegisterMakeShortJSONHandler() {
	h.internalEcho.POST(MakeShortJSONPath, func(c echo.Context) error {

		var shortRequest model.ShortURLRequest
		dec := json.NewDecoder(c.Request().Body)
		err := dec.Decode(&shortRequest)
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}

		userID := ""
		if h.HasUserID(c) {
			userID, err = h.GetUserID(c)
			if err != nil {
				return c.String(http.StatusUnauthorized, "unauthorized")
			}
		} else {
			userID = h.GenerateUserID(c)
		}

		shortUID, err := h.internalURLService.TryMakeShort(c.Request().Context(), userID, shortRequest.OriginalURL)
		var originalURLConflictError *service.OriginalURLConflictServiceError
		if err != nil && !errors.As(err, &originalURLConflictError) {
			return c.String(http.StatusInternalServerError, err.Error())
		}

		shortURL, err := url.JoinPath(h.internalBaseURL, shortUID)
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}

		c.Response().Header().Set(consts.HeaderContentType, consts.HeaderContentTypeJSON)
		h.SetUserID(c, userID)

		shortResponse := model.ShortURLResponse{
			ShortURL: shortURL,
		}

		if originalURLConflictError != nil {
			return c.JSON(http.StatusConflict, shortResponse)
		}

		return c.JSON(http.StatusCreated, shortResponse)
	})
}
