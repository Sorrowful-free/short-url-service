package handler

import (
	"fmt"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
)

func RegisterMakeShortHandler(e *echo.Echo) {
	e.POST("/", makeShortHandlerInternal)
}

func makeShortHandlerInternal(c echo.Context) error {

	originalURL, err := io.ReadAll(c.Request().Body)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	shortUID, err := internalURLService.TryMakeShort(string(originalURL))
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	shortURL := fmt.Sprintf("http://%s/%s", internalBaseURL, shortUID)
	c.Response().Header().Set("Content-Type", "text/plain")
	c.Response().WriteHeader(http.StatusCreated)

	fmt.Printf("process request for original URL:%s, with result:%s\n", originalURL, shortURL)
	return c.String(http.StatusCreated, shortURL)

}
