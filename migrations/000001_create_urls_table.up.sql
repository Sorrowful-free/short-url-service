CREATE TABLE short_urls (
    id SERIAL PRIMARY KEY,
    short_uid VARCHAR(16) NOT NULL,
    original_url VARCHAR(255) NOT NULL,
    UNIQUE(short_uid)
);

CREATE INDEX idx_short_uid ON short_urls(short_uid);

