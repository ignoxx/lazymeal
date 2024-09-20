-- +goose Up
CREATE TABLE IF NOT EXISTS meal_guide_whitelists (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	email TEXT UNIQUE NOT NULL,
	discount REAL NOT NULL
);

-- +goose Down
DROP TABLE IF EXISTS meal_guide_whitelists;
