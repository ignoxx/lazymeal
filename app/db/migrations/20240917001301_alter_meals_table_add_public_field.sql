-- +goose Up
ALTER TABLE meals ADD COLUMN isPublic BOOLEAN NOT NULL DEFAULT TRUE;

-- +goose Down
ALTER TABLE meals DROP COLUMN isPublic;
