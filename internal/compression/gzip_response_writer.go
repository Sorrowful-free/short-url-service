package compression

import (
	"compress/gzip"
	"net/http"

	"github.com/labstack/echo/v4"
)

type GzipResponseWriter struct {
	responseWriter http.ResponseWriter
	gzipWriter     *gzip.Writer
}

func NewGzipResponseWriter(response *echo.Response) *GzipResponseWriter {
	gzipWriter := gzip.NewWriter(response.Writer)
	return &GzipResponseWriter{
		responseWriter: response.Writer,
		gzipWriter:     gzipWriter,
	}
}

func (gzrw *GzipResponseWriter) Header() http.Header {
	return gzrw.responseWriter.Header()
}

func (gzrw *GzipResponseWriter) Write(p []byte) (int, error) {
	return gzrw.gzipWriter.Write(p)
}

func (gzrw *GzipResponseWriter) WriteHeader(statusCode int) {
	gzrw.responseWriter.WriteHeader(statusCode)
}

func (gzrw *GzipResponseWriter) Close() error {
	return gzrw.gzipWriter.Close()
}
