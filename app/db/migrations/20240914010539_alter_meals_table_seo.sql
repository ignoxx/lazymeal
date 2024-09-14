-- +goose Up
-- add columns: image_alt
ALTER TABLE meals
ADD COLUMN image_alt TEXT;

-- +goose Down
drop columns: image_alt
