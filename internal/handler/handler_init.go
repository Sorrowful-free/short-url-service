package handler

import (
	"fmt"
	"net/http"

	"github.com/Sorrowful-free/short-url-service/internal/consts"
	"github.com/Sorrowful-free/short-url-service/internal/crypto"
	"github.com/Sorrowful-free/short-url-service/internal/service"
	"github.com/labstack/echo/v4"
)

const (
	MakeShortPath          = "/"
	MakeShortJSONPath      = "/api/shorten"
	MakeShortBatchJSONPath = "/api/shorten/batch"
	MakeOriginalPath       = "/:id"
	OriginalPathParam      = "id"
	PingDBPath             = "/ping"
	GetUserURLsPath        = "/api/user/urls"

	FallbackUserID   = "0000000000000000"
	UserIDCookieName = "userID"
)

type Handlers struct {
	internalEcho            *echo.Echo
	internalURLService      service.ShortURLService
	internalUserIDEncryptor crypto.UserIDEncryptor
	internalBaseURL         string
}

func NewHandlers(echo *echo.Echo, urlService service.ShortURLService, baseURL string, userIDCriptoKey string) (*Handlers, error) {

	userIDEncryptor, err := crypto.NewSha256UserIDEncryptor(userIDCriptoKey)
	if err != nil {
		return nil, fmt.Errorf("failed to create user ID encryptor: %w", err)
	}
	return &Handlers{
		internalEcho:            echo,
		internalURLService:      urlService,
		internalBaseURL:         baseURL,
		internalUserIDEncryptor: userIDEncryptor,
	}, nil
}

func (h *Handlers) RegisterHandlers() *Handlers {
	h.RegisterMakeShortHandler()
	h.RegisterMakeOriginalHandler()
	h.RegisterMakeShortJSONHandler()
	h.RegisterMakeShortBatchJSONHandler()
	h.RegisterGetUserUrlsHandler()
	h.RegisterPingDBHandler()
	return h
}

func (h *Handlers) GenerateUserID(c echo.Context) string {
	userID, err := crypto.GenerateRandomSequenceString(consts.TestUserIDLength)
	if err != nil {
		return FallbackUserID
	}
	return userID
}

func (h *Handlers) GetUserID(c echo.Context) (string, error) {
	userIDCookie := c.Request().Cookies()[0].Value
	userID, err := h.internalUserIDEncryptor.Decrypt(userIDCookie)
	if err != nil {
		return FallbackUserID, err
	}
	return userID, nil
}

func (h *Handlers) SetUserID(c echo.Context, userID string) {
	if userID == "" {
		return
	}

	encryptedUserID, err := h.internalUserIDEncryptor.Encrypt(userID)
	if err != nil {
		return
	}

	c.SetCookie(&http.Cookie{
		Name:  UserIDCookieName,
		Value: encryptedUserID,
	})
}

func (h *Handlers) HasUserID(c echo.Context) bool {
	cookies := c.Request().Cookies()
	if len(cookies) == 0 {
		return false
	}

	for _, cookie := range cookies {
		if cookie.Name == UserIDCookieName {
			return true
		}
	}
	return false
}
