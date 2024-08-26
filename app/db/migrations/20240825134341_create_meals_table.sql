-- +goose Up
CREATE TABLE if not EXISTS meals(
	id integer PRIMARY KEY autoincrement,
	name text not null,
	category text not null,
	description text not null,
	short_description text not null,
	image_url text not null,
	image_preview text not null,
	calories integer not null,
	likes integer not null,
	total_effort integer not null,
	instruction_steps text not null,
	cook_time integer not null,
	prep_time integer not null,
	total_time integer not null,
	created_at datetime not null,
	updated_at datetime not null
);

-- +goose Down
DROP TABLE IF EXISTS meals;
