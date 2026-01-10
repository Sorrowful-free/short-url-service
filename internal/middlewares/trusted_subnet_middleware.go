package middlewares

import (
	"net"
	"net/http"
	"strings"

	"github.com/Sorrowful-free/short-url-service/internal/config"
	"github.com/labstack/echo/v4"
)

// TrustedSubnetMiddleware creates a middleware to verify that the request comes from a trusted subnet.
// The middleware checks the client's IP address and compares it with the CIDR from the configuration.
// If the IP is not in the trusted subnet, it returns status 403 Forbidden.
// If TrustedSubnet is not set in the config, the middleware allows all requests.
// Parameters:
//   - cfg: configuration with TrustedSubnet parameter
//
// Returns Echo middleware function.
func TrustedSubnetMiddleware(cfg *config.LocalConfig) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// If TrustedSubnet is not set, allow all requests
			if cfg.TrustedSubnet == "" {
				return next(c)
			}

			// Get the client's IP address
			ok, clientIP := tryGetClientIP(c)
			if !ok {
				return c.String(http.StatusForbidden, "forbidden")
			}

			// Parse CIDR from config
			_, ipNet, err := net.ParseCIDR(cfg.TrustedSubnet)
			if err != nil {
				// If CIDR is invalid, block all requests
				return c.String(http.StatusForbidden, "forbidden")
			}

			// Check if the IP is in the subnet
			if !ipNet.Contains(clientIP) {
				return c.String(http.StatusForbidden, "forbidden")
			}

			return next(c)
		}
	}
}

// tryGetClientIP extracts the client's IP address from the X-Real-IP header.
// Parses the value using ParseCIDR to support both IP and CIDR notation.
// Returns a tuple of (success flag, client IP address).
// The success flag is true if the IP was successfully extracted and parsed, false otherwise.
func tryGetClientIP(c echo.Context) (bool, net.IP) {
	realIP := c.Request().Header.Get("X-Real-IP")
	if realIP == "" {
		return false, nil
	}

	// X-Real-IP should contain one IP, but just in case take the first one
	ipStr := strings.TrimSpace(strings.Split(realIP, ",")[0])
	if ipStr == "" {
		return false, nil
	}

	// Try to parse as CIDR first
	ip, _, err := net.ParseCIDR(ipStr)
	if err == nil {
		return true, ip
	}

	// If ParseCIDR failed, try to parse as plain IP
	ip = net.ParseIP(ipStr)
	if ip != nil {
		return true, ip
	}

	return false, nil
}
