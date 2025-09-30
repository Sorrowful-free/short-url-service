package handler

import (
	"net/http"

	"github.com/Sorrowful-free/short-url-service/internal/consts"
	"github.com/labstack/echo/v4"
)

func (h *Handlers) RegisterMakeOriginalHandler() {
	h.internalEcho.GET(MakeOriginalPath, func(c echo.Context) error {
		shortUID := c.Param(OriginalPathParam)
		dto, err := h.internalURLService.TryMakeOriginal(c.Request().Context(), shortUID)
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}

		if dto.IsDeleted {
			return c.String(http.StatusGone, "short url is deleted")
		}

		c.Response().Header().Set(consts.HeaderContentType, consts.HeaderContentTypeText)
		c.Response().Header().Set(consts.HeaderLocation, dto.OriginalURL)

		return c.Redirect(http.StatusTemporaryRedirect, dto.OriginalURL)
	})
}
