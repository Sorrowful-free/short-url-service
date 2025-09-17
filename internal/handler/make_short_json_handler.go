package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"

	"github.com/Sorrowful-free/short-url-service/internal/consts"
	"github.com/Sorrowful-free/short-url-service/internal/model"
	"github.com/Sorrowful-free/short-url-service/internal/service/service_errors"
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
		shortUID, err := h.internalURLService.TryMakeShort(c.Request().Context(), shortRequest.OriginalURL)
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}

		shortURL, err := url.JoinPath(h.internalBaseURL, shortUID)
		var originalURLConflictError *service_errors.OriginalURLConflictServiceError
		if err != nil && !errors.As(err, &originalURLConflictError) {
			return c.String(http.StatusInternalServerError, err.Error())
		}

		c.Response().Header().Set(consts.HeaderContentType, consts.HeaderContentTypeJSON)
		c.Response().WriteHeader(http.StatusCreated)

		shortResponse := model.ShortURLResponse{
			ShortURL: shortURL,
		}

		if originalURLConflictError != nil {
			c.Response().Status = http.StatusConflict
		}

		enc := json.NewEncoder(c.Response().Writer)
		err = enc.Encode(shortResponse)

		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}

		return nil
	})
}
