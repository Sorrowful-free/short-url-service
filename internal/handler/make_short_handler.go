package handler

import (
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/Sorrowful-free/short-url-service/internal/consts"
	"github.com/Sorrowful-free/short-url-service/internal/service"
	"github.com/labstack/echo/v4"
)

func (h *Handlers) RegisterMakeShortHandler() {
	h.internalEcho.POST(MakeShortPath, func(c echo.Context) error {
		originalURL, err := io.ReadAll(c.Request().Body)
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

		var originalURLConflictError *service.OriginalURLConflictServiceError
		shortUID, err := h.internalURLService.TryMakeShort(c.Request().Context(), userID, string(originalURL))
		if err != nil && !errors.As(err, &originalURLConflictError) {
			return c.String(http.StatusInternalServerError, err.Error())
		}

		shortURL, err := url.JoinPath(h.internalBaseURL, shortUID)
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}

		c.Response().Header().Set(consts.HeaderContentType, consts.HeaderContentTypeText)
		h.SetUserID(c, userID)

		if originalURLConflictError != nil {
			return c.String(http.StatusConflict, shortURL)
		}

		return c.String(http.StatusCreated, shortURL)
	})
}
