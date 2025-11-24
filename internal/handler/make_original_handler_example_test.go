package handler_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/Sorrowful-free/short-url-service/internal/config"
	"github.com/Sorrowful-free/short-url-service/internal/consts"
	"github.com/Sorrowful-free/short-url-service/internal/handler"
	"github.com/Sorrowful-free/short-url-service/internal/service"
	"github.com/labstack/echo/v4"
)

func ExampleHandlers_RegisterMakeOriginalHandler() {
	// Создаем экземпляр Echo
	e := echo.New()

	// Создаем мок сервиса (в реальном приложении используется реальный сервис)
	// Для примера используем nil, в реальности нужен реализованный сервис
	var urlService service.ShortURLService

	// Получаем конфигурацию
	config := config.GetLocalConfig()

	// Создаем хэндлеры
	handlers, _ := handler.NewHandlers(e, "http://localhost:8080", urlService, config)
	handlers.RegisterHandlers()

	// Пример: GET запрос для получения оригинального URL по короткому идентификатору
	req := httptest.NewRequest(http.MethodGet, "/abc123", nil)
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	fmt.Printf("Status: %d\n", rec.Code)
	fmt.Printf("Location: %s\n", rec.Header().Get(consts.HeaderLocation))

	// Output:
	// Status: 307
	// Location: https://example.com/original-url
}

func ExampleHandlers_RegisterMakeOriginalHandler_deleted() {
	e := echo.New()
	var urlService service.ShortURLService
	config := config.GetLocalConfig()

	handlers, _ := handler.NewHandlers(e, "http://localhost:8080", urlService, config)
	handlers.RegisterHandlers()

	// Пример: запрос к удаленному URL
	req := httptest.NewRequest(http.MethodGet, "/deleted123", nil)
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	fmt.Printf("Status: %d\n", rec.Code)
	fmt.Printf("Body: %s\n", strings.TrimSpace(rec.Body.String()))

	// Output:
	// Status: 410
	// Body: short url is deleted
}

