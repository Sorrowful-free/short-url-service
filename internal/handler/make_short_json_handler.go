package handler

import (
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/Sorrowful-free/short-url-service/internal/consts"
	"github.com/Sorrowful-free/short-url-service/internal/model"
	"github.com/labstack/echo/v4"
)

func (h *Handlers) RegisterMakeShortJSONHandler() {
	h.internalEcho.POST(MakeShortJSONPath, func(c echo.Context) error {

		var shortRequest model.ShortRequest
		dec := json.NewDecoder(c.Request().Body)
		err := dec.Decode(&shortRequest)
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}
		shortUID, err := h.internalURLService.TryMakeShort(shortRequest.OriginalURL)
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}

		shortURL, err := url.JoinPath(h.internalBaseURL, shortUID)
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}

		c.Response().Header().Set(consts.HeaderContentType, consts.HeaderContentTypeJSON)
		c.Response().WriteHeader(http.StatusCreated)

		shortResponse := model.ShortResponse{
			ShortURL: shortURL,
		}

		enc := json.NewEncoder(c.Response().Writer)
		err = enc.Encode(shortResponse)
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}

		return nil
	})
}
