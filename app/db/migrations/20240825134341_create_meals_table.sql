-- +goose Up
CREATE TABLE IF NOT EXISTS meals (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    category TEXT NOT NULL,
    servings INTEGER NOT NULL DEFAULT 1, -- Number of servings (people)
    description TEXT NOT NULL,
    light_version_instructions TEXT, -- Instructions for simplifying the meal
    instructions TEXT NOT NULL, -- Full instructions for the meal
    image_url TEXT NOT NULL,
    calories INTEGER NOT NULL,
    protein INTEGER NOT NULL,
    cook_time INTEGER NOT NULL, -- in minutes
    prep_time INTEGER NOT NULL, -- in minutes
    total_time INTEGER NOT NULL, -- in minutes
    washing_effort INTEGER NOT NULL, -- 1 to 10
    peeling_effort INTEGER NOT NULL, -- 1 to 10
    cutting_effort INTEGER NOT NULL, -- 1 to 10
    items_required TEXT NOT NULL, -- Comma-separated list of required items
    ingredients TEXT NOT NULL, -- Comma-separated list of ingredients with amounts
    total_effort INTEGER NOT NULL, -- 1 to 10
    likes INTEGER NOT NULL DEFAULT 0,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- +goose Down
DROP TABLE IF EXISTS meals;
