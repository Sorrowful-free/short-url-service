package compression

import (
	"compress/gzip"
	"net/http"
)

type GzipRequestReader struct {
	request    *http.Request
	gzipReader *gzip.Reader
}

func NewGzipRequestReader(request *http.Request) (*GzipRequestReader, error) {
	gzipReader, err := gzip.NewReader(request.Body)
	if err != nil {
		return nil, err
	}
	return &GzipRequestReader{
		request:    request,
		gzipReader: gzipReader,
	}, nil
}

func (gzrr *GzipRequestReader) Header() http.Header {
	return gzrr.request.Header
}

func (gzrr *GzipRequestReader) Read(p []byte) (int, error) {
	return gzrr.gzipReader.Read(p)
}

func (gzrr *GzipRequestReader) Close() error {
	return gzrr.gzipReader.Close()
}
