package middlewares

import (
	"bytes"
	"encoding/json"
	"net/http"
	"os"
	"time"

	"github.com/Sorrowful-free/short-url-service/internal/config"
	"github.com/labstack/echo/v4"
)

type AuditAction string

const (
	AuditShorten AuditAction = "shorten"
	AuditFollow  AuditAction = "follow"
)

type AuditEvent struct {
	Timestamp   time.Time   `json:"ts"`
	Action      AuditAction `json:"action"`
	UserID      string      `json:"user_id"`
	OriginalURL string      `json:"url"`
}

func AuditMiddleWare(config *config.LocalConfig, action AuditAction) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if config.HasAuditFilePath() {
				err := AuditToFile(config.AuditFilePath, action)(c)
				if err != nil {
					return c.String(http.StatusInternalServerError, err.Error())
				}
			}
			if config.HasAuditURL() {
				err := AuditToURL(config.AuditURL, action)(c)
				if err != nil {
					return c.String(http.StatusInternalServerError, err.Error())
				}
			}
			return next(c)
		}
	}
}

func AuditToFile(filePath string, action AuditAction) echo.HandlerFunc {
	return func(c echo.Context) error {

		auditEvent := makeAuditEvent(c, action)
		auditEventJSON, err := json.Marshal(auditEvent)
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}
		err = os.WriteFile(filePath, auditEventJSON, os.ModeAppend)
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}
		return nil
	}
}

func AuditToURL(url string, action AuditAction) echo.HandlerFunc {
	return func(c echo.Context) error {

		auditEvent := makeAuditEvent(c, action)
		auditEventJSON, err := json.Marshal(auditEvent)
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}
		resp, err := http.Post(url, "application/json", bytes.NewBuffer(auditEventJSON))
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}
		resp.Body.Close()
		return nil
	}
}

func makeAuditEvent(c echo.Context, action AuditAction) AuditEvent {
	return AuditEvent{
		Timestamp:   time.Now(),
		Action:      action,
		UserID:      TryGetUserID(c),
		OriginalURL: c.Request().URL.Path,
	}
}
