ALTER TABLE short_urls ADD COLUMN user_id VARCHAR(16) NOT NULL;
CREATE INDEX idx_user_id ON short_urls(user_id);