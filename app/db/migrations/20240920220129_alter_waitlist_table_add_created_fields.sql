-- +goose Up
ALTER TABLE meal_guide_whitelists ADD COLUMN created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP;

-- +goose Down
ALTER TABLE meal_guide_whitelists DROP COLUMN created_at;
