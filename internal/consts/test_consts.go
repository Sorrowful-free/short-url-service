package consts

const (
	TestFileStoragePath     = "/test_data/short_urls_test.json"
	TestBaseURL             = "http://localhost:8080"
	TestUIDLength           = 8
	TestUIDRetryCount       = 10
	TestListenAddr          = "localhost:8080"
	TestOriginalURL         = "http://example.com"
	TestShortURL            = "http://localhost:8080/1234567890"
	TestShortUID            = "1234567890"
	TestOriginalURL2        = "http://example2.com"
	TestShortURL2           = "http://localhost:8080/2345678901"
	TestShortUID2           = "2345678901"
	TestDatabaseDSN         = "postgresql://postgres:postgres@localhost:5432?sslmode=disable"
	TestInvalideDatabaseDSN = "postgresql://postgres:postgres@postgres:4325/short_urls?sslmode=disable"
	TestUserID              = "1234567890"
	TestInvalidUserID       = "invalid user ID"
	TestUserIDLength        = 8
	TestUserIDCriptoKey     = "supersecretkey"
)
