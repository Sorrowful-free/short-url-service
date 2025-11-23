DB_URI = postgresql://postgres:postgres@postgres:5432/praktikum
MIGRATIONS_DIR = migrations
GO_BIN = go
PSQL_BIN = psql

migrate-up:
	$(PSQL_BIN) $(DB_URI) -f $(MIGRATIONS_DIR)/000001_create_urls_table.up.sql
	$(PSQL_BIN) $(DB_URI) -f $(MIGRATIONS_DIR)/000002_add_index_for_original_url.up.sql
	$(PSQL_BIN) $(DB_URI) -f $(MIGRATIONS_DIR)/000003_add_user_id_for_urls.up.sql
	$(PSQL_BIN) $(DB_URI) -f $(MIGRATIONS_DIR)/000004_add_is_deleted_field_to_short_urls.up.sql

migrate-down:
	$(PSQL_BIN) $(DB_URI) -f $(MIGRATIONS_DIR)/000001_create_urls_table.down.sql
	$(PSQL_BIN) $(DB_URI) -f $(MIGRATIONS_DIR)/000002_add_index_for_original_url.down.sql
	$(PSQL_BIN) $(DB_URI) -f $(MIGRATIONS_DIR)/000003_add_user_id_for_urls.down.sql
	$(PSQL_BIN) $(DB_URI) -f $(MIGRATIONS_DIR)/000004_add_is_deleted_field_to_short_urls.down.sql

migrate-reset:
	make migrate-down
	make migrate-up

gen_mocks:
	mockgen -source=internal/service/short_url_service.go -destination=mocks/mock_short_url_service.go -package=mocks
	mockgen -source=internal/repository/short_url_repository.go -destination=mocks/mock_short_url_repository.go -package=mocks
	
test:
	make gen_mocks
	$(GO_BIN) test -v ./...

run:
	make migrate-reset
	make test
	$(GO_BIN) run cmd/gophermart/main.go

build:
	make migrate-reset
	make test
	$(GO_BIN) build -o gophermart cmd/gophermart/main.go
