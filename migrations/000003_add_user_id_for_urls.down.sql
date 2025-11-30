DROP INDEX IF EXISTS idx_user_id;
ALTER TABLE short_urls DROP COLUMN user_id;
CREATE INDEX idx_user_id ON short_urls(user_id);