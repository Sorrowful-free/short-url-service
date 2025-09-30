package middlewares

import (
	"net/http"

	"github.com/Sorrowful-free/short-url-service/internal/consts"
	"github.com/Sorrowful-free/short-url-service/internal/crypto"
	"github.com/labstack/echo/v4"
)

const (
	FallbackUserID = "0000000000000000"
)

type SimpleAuthContext struct {
	echo.Context
	UserID string
}

func SimpleAuthMiddleware(userIDEncryptor crypto.UserIDEncryptor) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {

			userID := ""
			var err error
			if HasUserID(c) {
				userID, err = GetUserID(c, userIDEncryptor)
				if err != nil {
					return c.String(http.StatusUnauthorized, "unauthorized")
				}
			} else {
				userID = GenerateUserID()
			}

			SetUserID(c, userID, userIDEncryptor)

			return next(&SimpleAuthContext{
				Context: c,
				UserID:  userID,
			})
		}
	}
}

func GenerateUserID() string {
	userID, err := crypto.GenerateRandomSequenceString(consts.TestUserIDLength)
	if err != nil {
		return FallbackUserID
	}
	return userID
}

func GetUserID(c echo.Context, userIDEncryptor crypto.UserIDEncryptor) (string, error) {
	cookies := c.Request().Cookies()
	if len(cookies) == 0 {
		return FallbackUserID, nil
	}

	userIDCookie := ""
	for _, cookie := range cookies {
		if cookie.Name == consts.UserIDCookieName {
			userIDCookie = cookie.Value
			break
		}
	}

	if userIDCookie == "" {
		return FallbackUserID, nil
	}

	userID, err := userIDEncryptor.Decrypt(userIDCookie)
	if err != nil {
		return FallbackUserID, err
	}
	return userID, nil
}

func SetUserID(c echo.Context, userID string, userIDEncryptor crypto.UserIDEncryptor) {
	if userID == "" {
		return
	}

	encryptedUserID, err := userIDEncryptor.Encrypt(userID)
	if err != nil {
		return
	}

	c.SetCookie(&http.Cookie{
		Name:  consts.UserIDCookieName,
		Value: encryptedUserID,
	})
}

func HasUserID(c echo.Context) bool {
	cookies := c.Request().Cookies()
	if len(cookies) == 0 {
		return false
	}

	for _, cookie := range cookies {
		if cookie.Name == consts.UserIDCookieName {
			return true
		}
	}
	return false
}

func TryGetUserID(c echo.Context) string {
	if authContext, ok := c.(*SimpleAuthContext); ok {
		return authContext.UserID
	}
	return FallbackUserID
}
