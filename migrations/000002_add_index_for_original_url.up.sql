ALTER TABLE short_urls ADD CONSTRAINT original_url UNIQUE(original_url);
CREATE INDEX idx_original_url ON short_urls(original_url);