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

// SimpleAuthContext extends echo.Context with user identification.
// It provides access to the authenticated user ID in request handlers.
type SimpleAuthContext struct {
	echo.Context
	UserID string // The authenticated user ID
}

// SimpleAuthMiddleware creates an Echo middleware function that handles user authentication.
// It extracts or generates a user ID from cookies, encrypts it, and sets it in the response.
// If no user ID is found in cookies, a new one is generated.
// Parameters:
//   - userIDEncryptor: the encryptor used to encrypt/decrypt user IDs in cookies
//
// Returns an Echo middleware function.
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

// GenerateUserID generates a new random user ID.
// Returns a random string of TestUserIDLength characters, or FallbackUserID if generation fails.
func GenerateUserID() string {
	userID, err := crypto.GenerateRandomSequenceString(consts.TestUserIDLength)
	if err != nil {
		return FallbackUserID
	}
	return userID
}

// GetUserID extracts and decrypts the user ID from the request cookies.
// If no user ID cookie is found, it returns FallbackUserID.
// Parameters:
//   - c: the Echo context
//   - userIDEncryptor: the encryptor used to decrypt the user ID
//
// Returns the decrypted user ID and an error if decryption fails.
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

// SetUserID encrypts the user ID and sets it as a cookie in the response.
// If the user ID is empty or the cookie already exists with the same value, no action is taken.
// Parameters:
//   - c: the Echo context
//   - userID: the user ID to encrypt and set
//   - userIDEncryptor: the encryptor used to encrypt the user ID
func SetUserID(c echo.Context, userID string, userIDEncryptor crypto.UserIDEncryptor) {
	if userID == "" {
		return
	}

	encryptedUserID, err := userIDEncryptor.Encrypt(userID)
	if err != nil {
		return
	}

	cookies := c.Request().Cookies()
	for _, cookie := range cookies {
		if cookie.Name == consts.UserIDCookieName {
			if cookie.Value == encryptedUserID {
				return
			}
			break
		}
	}

	c.SetCookie(&http.Cookie{
		Name:  consts.UserIDCookieName,
		Value: encryptedUserID,
	})
}

// HasUserID checks if the request contains a user ID cookie.
// Returns true if a user ID cookie is present, false otherwise.
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

// TryGetUserID attempts to extract the user ID from the Echo context.
// If the context is a SimpleAuthContext, it returns the UserID field.
// Otherwise, it returns FallbackUserID.
// Returns the user ID from the context, or FallbackUserID if not available.
func TryGetUserID(c echo.Context) string {
	if authContext, ok := c.(*SimpleAuthContext); ok {
		return authContext.UserID
	}
	return FallbackUserID
}
