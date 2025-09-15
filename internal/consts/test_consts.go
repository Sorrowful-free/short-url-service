package consts

const (
	TestFileStoragePath     = "/test_data/short_urls_test.json"
	TestBaseURL             = "http://localhost:8080"
	TestUIDLength           = 8
	TestListenAddr          = "localhost:8080"
	TestOriginalURL         = "http://example.com"
	TestShortURL            = "http://localhost:8080/1234567890"
	TestShortUID            = "1234567890"
	TestDatabaseDSN         = "postgresql://postgres:postgres@localhost:5432?sslmode=disable"
	TestInvalideDatabaseDSN = "postgresql://postgres:postgres@postgres:4325/short_urls?sslmode=disable"
)
