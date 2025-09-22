package handler

import (
	"net/http"
	"net/url"

	"github.com/Sorrowful-free/short-url-service/internal/model"
	"github.com/labstack/echo/v4"
)

func (h *Handlers) RegisterGetUserUrlsHandler() {
	h.internalEcho.GET(GetUserURLsPath, func(c echo.Context) error {

		userID := ""
		var err error
		if h.HasUserID(c) {
			userID, err = h.GetUserID(c)
			if err != nil {
				return c.String(http.StatusUnauthorized, "unauthorized")
			}
		} else {
			userID = h.GenerateUserID(c)
		}

		shortURLDTOs, err := h.internalURLService.GetUserUrls(c.Request().Context(), userID)
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}

		if len(shortURLDTOs) == 0 {
			return c.String(http.StatusNoContent, "no content")
		}

		var getUserUrlsResponse model.UserShortURLResponse = make([]model.UserShortURLResponseDto, 0)
		for _, shortURLDTO := range shortURLDTOs {
			shortURL, err := url.JoinPath(h.internalBaseURL, shortURLDTO.ShortUID)
			if err != nil {
				return c.String(http.StatusInternalServerError, err.Error())
			}
			originalURL := shortURLDTO.OriginalURL
			getUserUrlsResponse = append(getUserUrlsResponse,
				model.UserShortURLResponseDto{ShortURL: shortURL, OriginalURL: originalURL})
		}
		return c.JSON(http.StatusOK, getUserUrlsResponse)
	})
}
