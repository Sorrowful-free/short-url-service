package handler_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/Sorrowful-free/short-url-service/internal/config"
	"github.com/Sorrowful-free/short-url-service/internal/handler"
	"github.com/Sorrowful-free/short-url-service/internal/service"
	"github.com/labstack/echo/v4"
)

func ExampleHandlers_RegisterPingDBHandler() {
	// Создаем экземпляр Echo
	e := echo.New()

	// Создаем простую реализацию сервиса для примера
	urlService := &service.ExampleService{PingError: false}

	// Получаем конфигурацию
	config := config.GetLocalConfig()

	// Создаем хэндлеры
	handlers, _ := handler.NewHandlers(e, "http://localhost:8080", urlService, config)
	handlers.RegisterHandlers()

	// Пример: GET запрос для проверки доступности базы данных
	req := httptest.NewRequest(http.MethodGet, "/ping", nil)
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	fmt.Printf("Status: %d\n", rec.Code)
	fmt.Printf("Body: %s\n", strings.TrimSpace(rec.Body.String()))

	// Output:
	// Status: 200
	// Body: pong
}

func ExampleHandlers_RegisterPingDBHandler_error() {
	e := echo.New()
	urlService := &service.ExampleService{PingError: true}
	config := config.GetLocalConfig()

	handlers, _ := handler.NewHandlers(e, "http://localhost:8080", urlService, config)
	handlers.RegisterHandlers()

	// Пример: запрос при недоступности базы данных
	req := httptest.NewRequest(http.MethodGet, "/ping", nil)
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	fmt.Printf("Status: %d\n", rec.Code)
	fmt.Printf("Body: %s\n", strings.TrimSpace(rec.Body.String()))

	// Output:
	// Status: 500
	// Body: database connection error
}
