package handler

import (
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/Sorrowful-free/short-url-service/internal/middlewares"
	"github.com/Sorrowful-free/short-url-service/internal/model"
	"github.com/labstack/echo/v4"
)

func (h *Handlers) RegisterMakeShortBatchJSONHandler() {
	h.internalEcho.POST(MakeShortBatchJSONPath, func(c echo.Context) error {
		var batchShortURLRequest model.BatchShortURLRequest

		dec := json.NewDecoder(c.Request().Body)
		err := dec.Decode(&batchShortURLRequest)
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}

		originalURLs := make([]string, len(batchShortURLRequest))
		for i, request := range batchShortURLRequest {
			originalURLs[i] = request.OriginalURL
		}

		authContext := c.(*middlewares.SimpleAuthContext)

		shortUIDs, err := h.internalURLService.TryMakeShortBatch(c.Request().Context(), authContext.UserID, originalURLs)
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}

		batchShortURLResponse := make([]model.BatchShortURLResponseDto, len(shortUIDs))
		for i, shortUID := range shortUIDs {
			shortURL, err := url.JoinPath(h.internalBaseURL, shortUID)
			if err != nil {
				return c.String(http.StatusInternalServerError, err.Error())
			}
			batchShortURLResponse[i] = model.BatchShortURLResponseDto{
				CorrelationID: batchShortURLRequest[i].CorrelationID,
				ShortURL:      shortURL,
			}
		}
		return c.JSON(http.StatusCreated, batchShortURLResponse)
	})
}
